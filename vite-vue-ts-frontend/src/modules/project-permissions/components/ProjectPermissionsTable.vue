<script setup lang="ts">
    import { h, computed } from 'vue';
    import { useI18n } from "vue-i18n";

    import { useDialog, NEmpty, NTooltip, NIcon } from 'naive-ui';
    import { IconTrash, IconEdit, IconEyeCheck, IconSquarePlus } from '@tabler/icons-vue';

    import { ProjectPermission } from '../models/project-permission.ts';
    import type { TableHeaderColumn } from '../../../shared/types/table-header-column';
    import { renderIcon } from '../../../shared/composables/naive-ui-icon';
    import ManageTable from '../../../shared/components/tables/ManageTable.vue';
    import ManageTableActionButtons from '../../../shared/components/tables/ManageTableActionButtons.vue';
    import RefreshAddActionsColumn from '../../../shared/components/tables/RefreshAddActionsColumn.vue';
    import AvatarUserName from '../../../shared/components/AvatarUserName.vue';
    import TextFilterInput from '../../../shared/components/TextFilterInput.vue';

    interface Props {
        loading: boolean;
        projectPermissions: ProjectPermission[];
    }

    const { t } = useI18n();
    const dialog = useDialog();

    const emit = defineEmits(['refresh', 'add', 'delete']);

    const props = defineProps<Props>();

    const columns = computed<TableHeaderColumn[]>(() => [
        {
            label: t("modules.projectPermission.components.projectPermissionsTable.header.columns.user"),
            field: "name",
            sortable: false,
        },
        {
            label: t("modules.projectPermission.components.projectPermissionsTable.header.columns.role"),
            field: "role",
            sortable: false,
        },
        {
            label: t("modules.projectPermission.components.projectPermissionsTable.header.columns.projectPermissions"),
            field: "projectPermissions",
            sortable: false,
            align: "center",
        },
        {
            label: t("modules.projectPermission.components.projectPermissionsTable.header.columns.taskPermissions"),
            field: "taskPermissions",
            sortable: false,
            align: "center",
        },
    ]);

    const userFilter = defineModel<string>("userFilter", {
        default: "",
    });
    const roleFilter = defineModel<string>("roleFilter", {
        default: "",
    });

    const onRefresh = () => {
        emit("refresh");
    };

    const onAdd = () => {
        emit("add");
    };

    const onConfirmDelete = (projectPermission: ProjectPermission, index: number) => {
        dialog.warning({
            title: t("modules.projectPermission.components.projectPermissionsTable.dialogs.deleteConfirmation.title"),
            icon: renderIcon(IconTrash)(24),
            content: () =>
                h('div', [
                    t("modules.projectPermission.components.projectPermissionsTable.dialogs.deleteConfirmation.message", { user: projectPermission.user.name, role: projectPermission.role.name }),
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
</script>

<template>
    <ManageTable size="small">
        <template #thead>
            <tr>
                <th v-for="column in columns" :key="column.field"
                    :class="{ 'doneo-text-center': column.align === 'center' }">
                    {{ column.label }}
                </th>
                <th class="doneo-table-actions-column">{{ t("shared.components.table.header.columns.actions") }}</th>
            </tr>
            <tr class="hide-mobile">
                <th>
                    <TextFilterInput clearable size="small"
                        :placeholder="t('modules.projectPermission.components.projectPermissionsTable.filters.user.placeholder')"
                        v-model:value="userFilter" />
                </th>
                <th>
                    <TextFilterInput clearable size="small"
                        :placeholder="t('modules.projectPermission.components.projectPermissionsTable.filters.role.placeholder')"
                        v-model:value="roleFilter" />
                </th>
                <th></th>
                <th></th>
                <th class="doneo-text-center">
                    <RefreshAddActionsColumn @refresh="onRefresh" @add="onAdd" />
                </th>
            </tr>
        </template>
        <template #tbody>
            <tr v-for="projectPermission, index in projectPermissions" :key="projectPermission.id">
                <td>
                    <AvatarUserName :user-id="projectPermission.user.id" :user-name="projectPermission.user.name" />
                </td>
                <td>{{ projectPermission.role.name }}</td>
                <td class="doneo-text-center">
                    <n-tooltip trigger="hover" v-if="projectPermission.role.permissions.allowUpdateProject">
                        <template #trigger>
                            <n-icon :size="22">
                                <IconEdit class="doneo-cursor-help" />
                            </n-icon>
                        </template>
                        {{
                            t("modules.projectPermission.components.projectPermissionsTable.body.columns.permissionsHints.updateProjectAllowed")
                        }}
                    </n-tooltip>
                    <n-icon :size="22" class="doneo-disabled-permission-icon" v-else>
                        <IconEdit />
                    </n-icon>
                    <n-tooltip trigger="hover" v-if="projectPermission.role.permissions.allowDeleteProject">
                        <template #trigger>
                            <n-icon :size="22">
                                <IconTrash class="doneo-cursor-help" />
                            </n-icon>
                        </template>
                        {{
                            t("modules.projectPermission.components.projectPermissionsTable.body.columns.permissionsHints.deleteProjectAllowed")
                        }}
                    </n-tooltip>
                    <n-icon :size="22" class="doneo-disabled-permission-icon" v-else>
                        <IconTrash />
                    </n-icon>
                    <n-tooltip trigger="hover" v-if="projectPermission.role.permissions.allowViewProject">
                        <template #trigger>
                            <n-icon :size="22">
                                <IconEyeCheck class="doneo-cursor-help" />
                            </n-icon>
                        </template>
                        {{
                            t("modules.projectPermission.components.projectPermissionsTable.body.columns.permissionsHints.viewProjectAllowed")
                        }}
                    </n-tooltip>
                    <n-icon :size="22" class="doneo-disabled-permission-icon" v-else>
                        <IconEyeCheck />
                    </n-icon>
                </td>
                <td class="doneo-text-center">
                    <n-tooltip trigger="hover" v-if="projectPermission.role.permissions.allowAddTask">
                        <template #trigger>
                            <n-icon :size="22">
                                <IconSquarePlus class="doneo-cursor-help" />
                            </n-icon>
                        </template>
                        {{
                            t("modules.projectPermission.components.projectPermissionsTable.body.columns.permissionsHints.addTaskAllowed")
                        }}
                    </n-tooltip>
                    <n-icon :size="22" class="doneo-disabled-permission-icon" v-else>
                        <IconSquarePlus />
                    </n-icon>
                    <n-tooltip trigger="hover" v-if="projectPermission.role.permissions.allowUpdateTask">
                        <template #trigger>
                            <n-icon :size="22">
                                <IconEdit class="doneo-cursor-help" />
                            </n-icon>
                        </template>
                        {{
                            t("modules.projectPermission.components.projectPermissionsTable.body.columns.permissionsHints.updateTaskAllowed")
                        }}
                    </n-tooltip>
                    <n-icon :size="22" class="doneo-disabled-permission-icon" v-else>
                        <IconEdit />
                    </n-icon>
                    <n-tooltip trigger="hover" v-if="projectPermission.role.permissions.allowDeleteTask">
                        <template #trigger>
                            <n-icon :size="22">
                                <IconTrash class="doneo-cursor-help" />
                            </n-icon>
                        </template>
                        {{
                            t("modules.projectPermission.components.projectPermissionsTable.body.columns.permissionsHints.deleteTaskAllowed")
                        }}
                    </n-tooltip>
                    <n-icon :size="22" class="doneo-disabled-permission-icon" v-else>
                        <IconTrash />
                    </n-icon>
                    <n-tooltip trigger="hover" v-if="projectPermission.role.permissions.allowViewTask">
                        <template #trigger>
                            <n-icon :size="22">
                                <IconEyeCheck class="doneo-cursor-help" />
                            </n-icon>
                        </template>
                        {{
                            t("modules.projectPermission.components.projectPermissionsTable.body.columns.permissionsHints.viewTaskAllowed")
                        }}
                    </n-tooltip>
                    <n-icon :size="22" class="doneo-disabled-permission-icon" v-else>
                        <IconEyeCheck />
                    </n-icon>
                </td>
                <td class="doneo-text-center">
                    <ManageTableActionButtons show-delete @delete="onConfirmDelete(projectPermission, index)" />
                </td>
            </tr>
            <tr>
                <td :colspan="columns.length + 1" v-if="projectPermissions.length < 1 && !props.loading">
                    <n-empty
                        :description="t('modules.projectPermission.components.projectPermissionsTable.warnings.noItemsFound')">
                    </n-empty>
                </td>
            </tr>
        </template>
    </ManageTable>
</template>

<style lang="css" scoped>
    @media (max-width: 768px) {
        .hide-mobile {
            display: none;
        }
    }

    .doneo-disabled-permission-icon {
        opacity: 0.1;
    }
</style>