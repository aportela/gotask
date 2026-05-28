<script setup lang="ts">
    import { onMounted, onBeforeUnmount, ref, reactive, shallowRef, watch } from 'vue';
    import { useI18n } from "vue-i18n";

    import { NModal, NCard } from 'naive-ui';

    import { type AjaxStateInterface, defaultAjaxState, defaultAjaxStateRunning } from '../../../shared/types/ajaxState';
    import { useLoadingStore } from '../../../stores/loading';
    import { useNotify } from '../../../shared/composables/notification';
    import { appBus } from '../../../shared/composables/bus';
    import { taskStatusService } from '../services/task-status';
    import { handleAPIError } from '../../../api/client/errorHandler';
    import type { TaskStatusResponse, SearchRequest } from '../types/dto';
    import { TaskStatus } from '../models/task-status';
    import TaskStatusesTable from '../components/TaskStatusesTable.vue';
    import TaskStatusForm from '../components/TaskStatusForm.vue';
    import { Sort } from '../../../shared/types/models/sort';
    import type { FormMode } from '../../../shared/types/form-mode';

    const { t } = useI18n();
    const { notify } = useNotify();

    const loadingStore = useLoadingStore();

    const state: AjaxStateInterface = reactive({ ...defaultAjaxState });

    const items = shallowRef<TaskStatus[]>([]);

    const sort = ref<Sort>(new Sort("name", "ASC"));

    const nameFilter = ref<string>("");

    const showForm = ref<boolean>(false);
    const formMode = ref<FormMode>("add");

    const selectedItem = ref<TaskStatus>(new TaskStatus());

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

    const onShowUpdateForm = (taskStatus: TaskStatus, _index: number) => {
        selectedItem.value = taskStatus;
        formMode.value = "update";
        showForm.value = true;
    };

    const onAdd = (taskStatus: TaskStatus) => {
        showForm.value = false;
        notify('success', t("modules.taskStatus.components.ManageTaskStatusesPage.notifications.taskStatusAdded", { name: taskStatus.name }));
        onRefresh();
    };

    const onUpdate = (taskStatus: TaskStatus) => {
        showForm.value = false;
        notify('success', t("modules.taskStatus.components.ManageTaskStatusesPage.notifications.taskStatusUpdated", { name: taskStatus.name }));
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
            const response = await taskStatusService.search(payload);
            items.value = response.taskStatuses.map((taskStatus: TaskStatusResponse) => new TaskStatus(taskStatus))
        } catch (error: unknown) {
            items.value = [];
            state.ajaxErrors = true;
            handleAPIError(error,
                (apiError) => {
                    switch (apiError.response?.status) {
                        case 401:
                            state.ajaxErrors = false;
                            appBus.emit({ type: "reauthRequired", payload: { emitter: "ManageTaskStatusesPage.onRefresh" } });
                            break;
                        default:
                            state.ajaxErrorMessage = t("modules.taskStatus.components.ManageTaskStatusesPage.errors.refreshError");
                            break;
                    }
                },
                (fatalError) => {
                    state.ajaxErrorMessage = t("modules.taskStatus.components.ManageTaskStatusesPage.errors.refreshError");
                    console.error("Unhandled API error", { file: "ManageTaskStatusesPage.vue", method: "onRefresh" }, { err: fatalError });
                });
        }
        finally {
            state.ajaxRunning = false;
            if (state.ajaxErrorMessage) {
                appBus.emit({ type: "remoteAPIError", payload: { errorMessage: state.ajaxErrorMessage } });
            }
        }
    };

    const onDelete = async (taskStatus: TaskStatus, _index?: number) => {
        if (taskStatus.id) {
            Object.assign(state, defaultAjaxStateRunning);
            try {
                await taskStatusService.delete(taskStatus.id);
                notify('success', t("modules.taskStatus.components.ManageTaskStatusesPage.notifications.taskStatusUpdated", { name: taskStatus.name }));
                onRefresh();
            } catch (error: unknown) {
                state.ajaxErrors = true;
                handleAPIError(error,
                    (apiError) => {
                        switch (apiError.response?.status) {
                            case 401:
                                state.ajaxErrors = false;
                                selectedItem.value = taskStatus;
                                appBus.emit({ type: "reauthRequired", payload: { emitter: "ManageTaskStatusesPage.onDelete" } });
                                break;
                            case 404:
                                state.ajaxErrorMessage = t("modules.taskStatus.components.ManageTaskStatusesPage.errors.notFoundError");
                                break;
                            default:
                                state.ajaxErrorMessage = t("modules.taskStatus.components.ManageTaskStatusesPage.errors.deleteError");
                                break;
                        }
                    },
                    (fatalError) => {
                        state.ajaxErrorMessage = t("modules.taskStatus.components.ManageTaskStatusesPage.errors.deleteError");
                        console.error("Unhandled API error", { file: "ManageTaskStatusesPage.vue", method: "onRefresh" }, { err: fatalError });
                    });
            } finally {
                state.ajaxRunning = false;
                if (state.ajaxErrorMessage) {
                    appBus.emit({ type: "remoteAPIError", payload: { errorMessage: state.ajaxErrorMessage } });
                }
            }
        } else {
            console.error("task status id not set", { file: "ManageTaskStatusesPage.vue", method: "onDelete" });
        }
    };

    let stopBusReauthListener: () => void;

    onMounted(() => {
        onRefresh();
        stopBusReauthListener = appBus.on("reauthValidNotify", async (payload) => {
            if (payload.to.includes("ManageTaskStatusesPage.onRefresh")) {
                onRefresh();
            } else if (payload.to.includes("ManageTaskStatusesPage.onDelete")) {
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
        <TaskStatusForm :mode="formMode == 'add' ? 'add' : 'update'" :taskStatusId="selectedItem.id" style="width: 40%;"
            @add="onAdd" @update="onUpdate" @cancel="onCancel" />
    </n-modal>

    <n-card :title="t('modules.taskStatus.components.ManageTaskStatusesPage.header.title')">
        <TaskStatusesTable :task-statuses="items" :loading="state.ajaxRunning" @refresh="onRefresh" @add="onShowAddForm"
            @update="onShowUpdateForm" @delete="onDelete" @textfilter-keydown-enter="onRefresh" :sort-field="sort.field"
            :sort-order="sort.order" @toggle-sort="onToggleSort" v-model:task-status-name-filter="nameFilter" />
    </n-card>
</template>

<style lang="css" scoped></style>