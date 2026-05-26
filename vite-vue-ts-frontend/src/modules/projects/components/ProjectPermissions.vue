<script setup lang="ts">
    import { ref, reactive, h, onMounted, onBeforeUnmount, watch, type CSSProperties } from "vue";
    import { useI18n } from "vue-i18n";

    import { NCard, useDialog } from "naive-ui";
    import { IconTrash } from "@tabler/icons-vue";

    import { type AjaxStateInterface, defaultAjaxState, defaultAjaxStateRunning } from '../../../shared/types/ajaxState';
    import { useLoadingStore } from '../../../stores/loading';

    import { appBus } from '../../../shared/composables/bus';

    import { renderIcon } from "../../../shared/composables/naive-ui-icon.ts";

    import { projectPermissionService } from "../../project-permissions/services/project-permission.ts";
    import { handleAPIError } from '../../../api/client/errorHandler';

    import AvatarUserName from "../../../shared/components/AvatarUserName.vue";

    import ManageTable from '../../../shared/components/tables/ManageTable.vue';
    import ManageTableActionButtons from '../../../shared/components/tables/ManageTableActionButtons.vue';


    import type { SearchResponse } from "../../project-permissions/types/dto.ts";
    import { ProjectPermission } from "../../project-permissions/models/project-permission.ts";

    interface ProjectPermissionsProps {
        style?: string | CSSProperties;
        projectId: string;
    }

    const props = defineProps<ProjectPermissionsProps>();

    const emit = defineEmits(['delete']);

    const { t } = useI18n();

    const loadingStore = useLoadingStore();

    const state: AjaxStateInterface = reactive({ ...defaultAjaxState });

    watch(state, (newValue: AjaxStateInterface) => {
        loadingStore.set(newValue.ajaxRunning);
    });

    const projectPermissions = ref<ProjectPermission[]>();

    const dialog = useDialog();

    const onRefresh = async () => {
        Object.assign(state, defaultAjaxStateRunning);
        try {
            const results: SearchResponse = await projectPermissionService.search(props.projectId);
            projectPermissions.value = results.projectPermissions.map((permission) => new ProjectPermission(permission));
        } catch (error: unknown) {
            state.ajaxErrors = true;
            handleAPIError(error,
                (apiError) => {
                    switch (apiError.response?.status) {
                        case 401:
                            state.ajaxErrors = false;
                            appBus.emit({ type: "reauthRequired", payload: { emitter: "ProjectPermissions.onRefresh" } });
                            break;
                        case 404:
                            state.ajaxErrorMessage = t("modules.projectPriority.components.ProjectPriorityForm.errors.notFoundError");
                            break;
                        default:
                            state.ajaxErrorMessage = t("modules.projectPriority.components.ProjectPriorityForm.errors.loadError");
                            break;
                    }
                },
                (fatalError) => {
                    state.ajaxErrorMessage = t("modules.projectPriority.components.ProjectPriorityForm.errors.loadError");
                    console.error("Unhandled API error", { file: "ProjectPermissions.vue", method: "onRefresh" }, { err: fatalError });
                });
        } finally {
            state.ajaxRunning = false;
        }
    };

    const onConfirmDelete = (projectPermission: ProjectPermission, index: number) => {
        dialog.warning({
            title: t("modules.projectPriority.components.ProjectPrioritiesTable.dialogs.deleteConfirmation.title"),
            icon: renderIcon(IconTrash)(24),
            content: () =>
                h('div', [
                    t("modules.projectPriority.components.ProjectPrioritiesTable.dialogs.deleteConfirmation.message", { name: projectPermission.user.name }),
                    h('br'),
                    h('br'),
                    t("shared.components.dialogs.confirmation.continueMessage"),
                ]),
            positiveText: t("shared.buttons.Delete.label"),
            negativeText: t("shared.buttons.Cancel.label"),
            onPositiveClick: () => {
                emit("delete", projectPermission, index)
            },
        });
    };
    let stopBusReauthListener: () => void;

    onMounted(() => {
        onRefresh();
        stopBusReauthListener = appBus.on("reauthValidNotify", async (payload) => {
            if (payload.to.includes("ProjectPermissions.onRefresh")) {
                onRefresh();
            }
        });
    });

    onBeforeUnmount(() => {
        stopBusReauthListener();
    });

</script>

<template>
    <n-card bordered :style="props.style">
        <ManageTable size="small">
            <template #thead>
                <tr>
                    <th>Name</th>
                    <th>Role</th>
                    <th>Permissions</th>
                    <th class="doneo-table-actions-column">{{
                        t("shared.components.table.header.columns.actions") }}</th>
                </tr>
            </template>
            <template #tbody>
                <tr v-for="projectPermission, index in projectPermissions" :key="projectPermission.id">
                    <td>
                        <AvatarUserName :user-id="projectPermission.user.id" :user-name="projectPermission.user.name" />
                    </td>
                    <td>{{ projectPermission.role.name }}</td>
                    <td>[1][2][3][4]</td>
                    <td class="doneo-text-center">
                        <ManageTableActionButtons show-delete @delete="onConfirmDelete(projectPermission, index)" />
                    </td>
                </tr>
            </template>
        </ManageTable>
    </n-card>
</template>

<style lang="css" scoped></style>