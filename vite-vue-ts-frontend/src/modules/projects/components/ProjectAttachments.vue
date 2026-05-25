<script setup lang="ts">
    import type { CSSProperties } from "vue";
    import { useI18n } from "vue-i18n";

    import { NCard, NAvatar } from "naive-ui";

    import ManageTable from '../../../shared/components/tables/ManageTable.vue';
    import ManageTableActionButtons from '../../../shared/components/tables/ManageTableActionButtons.vue';

    interface ProjectAttachmentsProps {
        style?: string | CSSProperties;
        projectId: string;
    }

    const props = defineProps<ProjectAttachmentsProps>();

    const { t } = useI18n();
</script>

<template>
    <n-card bordered :style="props.style">
        <ManageTable size="small">
            <template #thead>
                <tr>
                    <th>Filename</th>
                    <th>Size</th>
                    <th>Mime</th>
                    <th>Uploaded at</th>
                    <th>Uploaded by</th>
                    <th class="doneo-table-actions-column">{{
                        t("shared.components.table.header.columns.actions") }}</th>
                </tr>
            </template>
            <template #tbody>
                <tr v-for="index in [1, 2, 3, 4, 5]" :key="index">
                    <td>
                        test.pdf
                    </td>
                    <td>8230Kb</td>
                    <td>application/pdf</td>
                    <td>23/02/2026</td>
                    <td>
                        <div class="doneo-flex-center-align" style="gap: 8px;">
                            <n-avatar size="small" src="https://i.pravatar.cc/32?u=019e5b68-f701-7af7-a694-c8146e3366a1"
                                class="avatar" />
                            John Doe
                        </div>
                    </td>
                    <td class="doneo-text-center">
                        <ManageTableActionButtons show-delete show-download />
                    </td>
                </tr>
            </template>
        </ManageTable>
    </n-card>
</template>

<style lang="css" scoped></style>