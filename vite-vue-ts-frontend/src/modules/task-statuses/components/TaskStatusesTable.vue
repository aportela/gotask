<script setup lang="ts">
    import { h, computed } from 'vue';
    import { useI18n } from "vue-i18n";

    import { useDialog, NFlex, NEmpty, NTag } from 'naive-ui';
    import { IconTrash } from '@tabler/icons-vue';

    import { TaskStatus } from '../models/task-status';
    import type { TableHeaderColumn } from '../../../shared/types/table-header-column';
    import { type SortOrder } from '../../../shared/types/common';
    import { renderIcon } from '../../../shared/composables/naive-ui-icon';
    import ManageTable from '../../../shared/components/tables/ManageTable.vue';
    import TextFilterInput from '../../../shared/components/TextFilterInput.vue';
    import TableCellHeaderSortIcon from '../../../shared/components/tables/TableCellHeaderSortIcon.vue';
    import UpdateDeleteActionsColumn from '../../../shared/components/tables/UpdateDeleteActionsColumn.vue';
    import RefreshAddActionsColumn from '../../../shared/components/tables/RefreshAddActionsColumn.vue';
    import { getNaiveUITagColorProperty } from '../../../shared/composables/color';
    import RemoteAPIAlert from '../../../shared/components/alerts/RemoteAPIAlert.vue';

    interface Props {
        loading: boolean;
        taskStatuses: TaskStatus[];
        sortField: string;
        sortOrder: SortOrder;
        errorMessage?: string | null;
    }

    const { t } = useI18n();

    const emit = defineEmits(['refresh', 'add', 'update', 'delete', 'toggleSort', 'textfilterKeydownEnter']);

    const props = defineProps<Props>();

    const columns = computed<TableHeaderColumn[]>(() => [
        {
            label: t("modules.taskStatus.components.TaskStatusesTable.header.columns.name"),
            field: "name",
            sortable: true,
        }
    ]);

    const taskStatusNameFilter = defineModel<string>("taskStatusNameFilter", {
        default: "",
    });

    const dialog = useDialog();

    const onToggleSort = (field: string) => {
        emit("toggleSort", field);
    };

    const onRefresh = () => {
        emit("refresh");
    };

    const onAdd = () => {
        emit("add");
    };

    const onUpdate = (taskStatus: TaskStatus, index: number) => {
        emit("update", taskStatus, index);
    };

    const onConfirmDelete = (taskStatus: TaskStatus, index: number) => {
        dialog.warning({
            title: t("modules.taskStatus.components.TaskStatusesTable.dialogs.deleteConfirmation.title"),
            icon: renderIcon(IconTrash)(24),
            content: () =>
                h('div', [
                    t("modules.taskStatus.components.TaskStatusesTable.dialogs.deleteConfirmation.message", { name: taskStatus.name }),
                    h('br'),
                    h('br'),
                    t("shared.components.dialogs.confirmation.continueMessage"),
                ]),
            positiveText: t("shared.buttons.Delete.label"),
            negativeText: t("shared.buttons.Cancel.label"),
            onPositiveClick: () => {
                emit("delete", taskStatus, index)
            },
        });
    };

    const onTextFilterKeyDownEnter = () => {
        emit("textfilterKeydownEnter");
    };

</script>

<template>
    <ManageTable size="small">
        <template #thead>
            <tr class="doneo-table-header-click-action">
                <th v-for="column in columns" :key="column.field" @click="column.sortable && onToggleSort(column.field)"
                    :class="{ 'doneo-cursor-pointer': column.sortable, 'doneo-text-center': column.align === 'center' }">
                    <n-flex justify="space-between" v-if="column.sortable">
                        <span>{{ column.label }}</span>
                        <TableCellHeaderSortIcon v-if="props.sortField === column.field" :order="props.sortOrder" />
                    </n-flex>
                    <span v-else>{{ column.label }}</span>
                </th>
                <th class="doneo-table-actions-column">{{ t("shared.components.table.header.columns.actions") }}</th>
            </tr>
            <tr>
                <th>
                    <TextFilterInput clearable size="small"
                        :placeholder="t('modules.taskStatus.components.TaskStatusesTable.filters.name.placeholder')"
                        v-model:value="taskStatusNameFilter" @keydown-enter="onTextFilterKeyDownEnter" />
                </th>
                <th class="doneo-text-center">
                    <RefreshAddActionsColumn @refresh="onRefresh" @add="onAdd" />
                </th>
            </tr>
        </template>
        <template #tbody v-if="!props.errorMessage">
            <tr v-for="taskStatus, index in taskStatuses" :key="taskStatus.id ?? index">
                <td>
                    <n-tag :color="getNaiveUITagColorProperty(taskStatus.hexColor ?? '#888888')">{{ taskStatus.name
                        }}</n-tag>
                </td>
                <td class="doneo-text-center">
                    <UpdateDeleteActionsColumn @update="onUpdate(taskStatus, index)"
                        @delete="onConfirmDelete(taskStatus, index)" />
                </td>
            </tr>
            <tr>
                <td :colspan="columns.length + 1" v-if="taskStatuses.length < 1 && !props.loading">
                    <n-empty :description="t('modules.taskStatus.components.TaskStatusesTable.warnings.noItemsFound')">
                    </n-empty>
                </td>
            </tr>
        </template>
        <template #error v-else>
            <tr>
                <td :colspan="columns.length + 1" v-if="props.errorMessage && !props.loading">
                    <RemoteAPIAlert type="error" :title="t('shared.errorMessages.Error')"
                        :message="props.errorMessage" />
                </td>
            </tr>
        </template>
    </ManageTable>
</template>

<style lang="css" scoped></style>