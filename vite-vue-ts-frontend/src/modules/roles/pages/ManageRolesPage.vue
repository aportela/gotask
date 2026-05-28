<script setup lang="ts">
    import { onMounted, onBeforeUnmount, ref, reactive, shallowRef, watch } from 'vue';
    import { useI18n } from "vue-i18n";

    import { NModal, NCard } from 'naive-ui';

    import { type AjaxStateInterface, defaultAjaxState, defaultAjaxStateRunning } from '../../../shared/types/ajaxState';
    import { useLoadingStore } from '../../../stores/loading';
    import { useNotify } from '../../../shared/composables/notification';
    import { appBus } from '../../../shared/composables/bus';
    import { roleService as roleService } from '../services/role';
    import { handleAPIError } from '../../../api/client/errorHandler';
    import type { RoleResponse, SearchRequest } from '../types/dto';
    import { Role } from '../models/role';
    import RolesTable from '../components/RolesTable.vue';
    import RoleForm from '../components/RoleForm.vue';
    import { Sort } from '../../../shared/types/models/sort';
    import type { FormMode } from '../../../shared/types/form-mode';


    const { t } = useI18n();
    const { notify } = useNotify();

    const loadingStore = useLoadingStore();

    const state: AjaxStateInterface = reactive({ ...defaultAjaxState });

    const items = shallowRef<Role[]>([]);

    const sort = ref<Sort>(new Sort("name", "ASC"));

    const nameFilter = ref<string>("");

    const showForm = ref<boolean>(false);
    const formMode = ref<FormMode>("add");

    const selectedItem = ref<Role>(new Role());

    watch(state, (newValue: AjaxStateInterface) => {
        loadingStore.set(newValue.ajaxRunning);
    });

    const onToggleSort = (field: string) => {
        sort.value.toggleSort(field);
        onRefresh();
    };

    const onShowAddForm = () => {
        formMode.value = "add";
        showForm.value = true;
    };

    const onShowUpdateForm = (role: Role, _index?: number) => {
        selectedItem.value = role;
        formMode.value = "update";
        showForm.value = true;
    };

    const onAdd = (role: Role) => {
        showForm.value = false;
        notify('success', t("modules.role.components.ManageRolesPage.notifications.roleAdded", { name: role.name }));
        onRefresh();
    };

    const onUpdate = (role: Role) => {
        showForm.value = false;
        notify('success', t("modules.role.components.ManageRolesPage.notifications.roleUpdated", { name: role.name }));
        onRefresh();
    };

    const onCancel = () => {
        showForm.value = false;
    };

    const onRefresh = async () => {
        Object.assign(state, defaultAjaxStateRunning);
        try {
            const payload: SearchRequest = {
                pager: {
                    currentPage: 1,
                    resultsPage: 0,
                },
                order: {
                    field: sort.value.field,
                    sort: sort.value.order,
                },
                filter: {
                    name: nameFilter.value,
                }
            };
            const response = await roleService.search(payload);
            items.value = response.roles.map((role: RoleResponse) => new Role(role));
        } catch (error: unknown) {
            items.value = [];
            state.ajaxErrors = true;
            handleAPIError(error,
                (apiError) => {
                    switch (apiError.response?.status) {
                        case 401:
                            state.ajaxErrors = false;
                            appBus.emit({ type: "reauthRequired", payload: { emitter: "ManageRolesPage.onRefresh" } });
                            break;
                        default:
                            state.ajaxErrorMessage = t("modules.role.components.ManageRolesPage.errors.refreshError");
                            break;
                    }
                },
                (fatalError) => {
                    state.ajaxErrorMessage = t("modules.role.components.ManageRolesPage.errors.refreshError");
                    console.error("Unhandled API error", { file: "ManageRolesPage.vue", method: "onRefresh" }, { err: fatalError });
                });
        }
        finally {
            state.ajaxRunning = false;
        }
    };

    const onDelete = async (role: Role, _index?: number) => {
        if (role.id) {
            Object.assign(state, defaultAjaxStateRunning);
            try {
                await roleService.delete(role.id);
                notify('success', t("modules.role.components.ManageRolesPage.notifications.roleDeleted", { name: role.name }));
                onRefresh();
            } catch (error: unknown) {
                state.ajaxErrors = true;
                handleAPIError(error,
                    (apiError) => {
                        switch (apiError.response?.status) {
                            case 401:
                                state.ajaxErrors = false;
                                selectedItem.value = role;
                                appBus.emit({ type: "reauthRequired", payload: { emitter: "ManageRolesPage.onDelete" } });
                                break;
                            case 404:
                                state.ajaxErrorMessage = t("modules.role.components.ManageRolesPage.errors.notFoundError");
                                break;
                            default:
                                state.ajaxErrorMessage = t("modules.role.components.ManageRolesPage.errors.deleteError");
                                break;
                        }
                    },
                    (fatalError) => {
                        state.ajaxErrorMessage = t("modules.role.components.ManageRolesPage.errors.deleteError");
                        console.error("Unhandled API error", { file: "ManageRolesPage.vue", method: "onRefresh" }, { err: fatalError });
                    });
            } finally {
                state.ajaxRunning = false;
            }
        } else {
            console.error("role id not set", { file: "ManageRolesPage.vue", method: "onDelete" });
        }
    };

    let stopBusReauthListener: () => void;

    onMounted(() => {
        onRefresh();
        stopBusReauthListener = appBus.on("reauthValidNotify", async (payload) => {
            if (payload.to.includes("ManageRolesPage.onRefresh")) {
                onRefresh();
            } else if (payload.to.includes("ManageRolesPage.onDelete")) {
                onDelete(selectedItem.value);
            }
        });
    });

    onBeforeUnmount(() => {
        stopBusReauthListener();
    });
</script>

<template>
    <n-modal v-model:show="showForm">
        <RoleForm :mode="formMode == 'add' ? 'add' : 'update'" :role-id="selectedItem.id" style="width: 40%;"
            @add="onAdd" @update="onUpdate" @cancel="onCancel" />
    </n-modal>

    <n-card :title="t('modules.role.components.ManageRolesPage.header.title')">
        <RolesTable :roles="items" :loading="state.ajaxRunning" @refresh="onRefresh" @add="onShowAddForm"
            @update="onShowUpdateForm" @delete="onDelete" @textfilter-keydown-enter="onRefresh" :sort-field="sort.field"
            :sort-order="sort.order" @toggle-sort="onToggleSort" v-model:role-name-filter="nameFilter"
            :error-message="state.ajaxErrorMessage" />
    </n-card>
</template>

<style lang="css"></style>