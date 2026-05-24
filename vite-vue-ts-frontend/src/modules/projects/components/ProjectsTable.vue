<script setup lang="ts">
    import { h, computed } from 'vue';
    import { useI18n } from "vue-i18n";

    import { useDialog, NFlex, NEmpty, NTag, NAvatar } from 'naive-ui';
    import { IconTrash } from '@tabler/icons-vue';

    import { Project } from '../models/project';
    import type { TableHeaderColumn } from '../../../shared/types/table-header-column';
    import { type SortOrder } from '../../../shared/types/common';
    import { renderIcon } from '../../../shared/composables/naive-ui-icon';
    import ManageTable from '../../../shared/components/tables/ManageTable.vue';
    import TextFilterInput from '../../../shared/components/TextFilterInput.vue';
    import TableCellHeaderSortIcon from '../../../shared/components/tables/TableCellHeaderSortIcon.vue';
    import UpdateDeleteActionsColumn from '../../../shared/components/tables/UpdateDeleteActionsColumn.vue';
    import RefreshAddActionsColumn from '../../../shared/components/tables/RefreshAddActionsColumn.vue';
    import ProjectPrioritySelector from '../../project-priorities/components/ProjectPrioritySelector.vue';
    import DateFilterSelect from '../../../shared/components/selectors/DateFilterSelect.vue';
    import { getNaiveUITagColorProperty } from '../../../shared/composables/color';

    interface Props {
        loading: boolean;
        projects: Project[];
        sortField: string;
        sortOrder: SortOrder;
    }

    const { t } = useI18n();

    const emit = defineEmits(['refresh', 'add', 'update', 'delete', 'toggleSort', 'textfilterKeydownEnter']);

    const props = defineProps<Props>();

    const columns = computed<TableHeaderColumn[]>(() => [
        {
            label: "Key",
            field: "key",
            sortable: true,
        },
        {
            label: "Type",
            field: "type",
            sortable: true,
        },
        {
            label: "Priority",
            field: "priority",
            sortable: true,
        },
        {
            label: "Status",
            field: "status",
            sortable: true,
        },
        {
            label: "Summary",
            field: "summary",
            sortable: true,
        },
        {
            label: "Created At",
            field: "createdAt",
            sortable: true,
        },
        {
            label: "Created by",
            field: "createdBy",
            sortable: true,
        },
    ]);


    const projectKeyFilter = defineModel<string>("projectKeyFilter", {
        default: "",
    });

    const projectSummaryFilter = defineModel<string>("projectSummaryFilter", {
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

    const onUpdate = (project: Project, index: number) => {
        emit("update", project, index);
    };

    const onConfirmDelete = (project: Project, index: number) => {
        dialog.warning({
            title: t("modules.projectPriority.components.ProjectPrioritiesTable.dialogs.deleteConfirmation.title"),
            icon: renderIcon(IconTrash)(24),
            content: () =>
                h('div', [
                    t("modules.projectPriority.components.ProjectPrioritiesTable.dialogs.deleteConfirmation.message", { name: project.summary }),
                    h('br'),
                    h('br'),
                    t("shared.components.dialogs.confirmation.continueMessage"),
                ]),
            positiveText: t("shared.buttons.Delete.label"),
            negativeText: t("shared.buttons.Cancel.label"),
            onPositiveClick: () => {
                emit("delete", project, index)
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
            <tr class="table-header-click-action">
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
            <tr class="hide-mobile">
                <th>
                    <TextFilterInput clearable size="small"
                        :placeholder="t('modules.projectPriority.components.ProjectPrioritiesTable.filters.name.placeholder')"
                        v-model:value="projectKeyFilter" @keydown-enter="onTextFilterKeyDownEnter" />
                </th>
                <th>
                    <ProjectPrioritySelector />
                </th>
                <th>
                    <ProjectPrioritySelector />
                </th>
                <th>
                    <ProjectPrioritySelector />
                </th>
                <th>
                    <TextFilterInput clearable size="small"
                        :placeholder="t('modules.projectPriority.components.ProjectPrioritiesTable.filters.name.placeholder')"
                        v-model:value="projectSummaryFilter" @keydown-enter="onTextFilterKeyDownEnter" />
                </th>
                <th>
                    <DateFilterSelect />
                </th>
                <th></th>
                <th class="doneo-text-center">
                    <RefreshAddActionsColumn @refresh="onRefresh" @add="onAdd" />
                </th>
            </tr>
        </template>
        <template #tbody>
            <tr v-for="project, index in projects" :key="project.id">
                <td>{{ project.key }}</td>
                <td><n-tag type="info" :bordered="false" :color="getNaiveUITagColorProperty(project.type.hexColor)">{{
                    project.type.name }}</n-tag></td>
                <td><n-tag type="info" :bordered="false"
                        :color="getNaiveUITagColorProperty(project.priority.hexColor)">{{ project.priority.name
                        }}</n-tag></td>
                <td><n-tag type="info" :bordered="false" :color="getNaiveUITagColorProperty(project.status.hexColor)">{{
                    project.status.name }}</n-tag></td>
                <td>{{ project.summary }}</td>
                <td>{{ project.createdAt.toLocaleString() }}</td>
                <td>
                    <div class="doneo-flex-center-align" style="gap: 8px;">
                        <n-avatar v-if="project.createdBy.avatarUrl" :src="project.createdBy.avatarUrl"
                            class="avatar" />
                        {{ project.createdBy.name }}
                    </div>

                </td>
                <td class="doneo-text-center">
                    <UpdateDeleteActionsColumn @update="onUpdate(project, index)"
                        @delete="onConfirmDelete(project, index)" />
                </td>
            </tr>
            <tr>
                <td :colspan="columns.length + 1" v-if="projects.length < 1 && !props.loading">
                    <n-empty
                        :description="t('modules.projectPriority.components.ProjectPrioritiesTable.warnings.noItemsFound')">
                    </n-empty>
                </td>
            </tr>
        </template>
    </ManageTable>
</template>

<style lang="css" scoped>

    .table-header-click-action th:not(:last-of-type) .n-icon {
        margin-top: 4px;
    }

    @media (max-width: 768px) {
        .hide-mobile {
            display: none;
        }
    }
</style>