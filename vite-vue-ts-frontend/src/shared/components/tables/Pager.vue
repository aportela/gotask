<script setup lang="ts">
    import { computed } from "vue";
    import { useI18n } from "vue-i18n";

    import { NFlex, NPagination } from 'naive-ui';
    import type { PaginationSizeOption } from "naive-ui";

    interface PagerProps {
        totalResults: number;
        totalPages: number;
    };

    const props = defineProps<PagerProps>();

    const { t } = useI18n();

    const pageSizes = computed<PaginationSizeOption[]>(() => [
        {
            label: t("All results"),
            value: 0
        },
        {
            label: t("labelSelectorResultsPage", { number: 5 }),
            value: 5
        },
        {
            label: t("labelSelectorResultsPage", { number: 10 }),
            value: 10
        },
        {
            label: t("labelSelectorResultsPage", { number: 20 }),
            value: 20
        },
        {
            label: t("labelSelectorResultsPage", { number: 50 }),
            value: 50
        },
        {
            label: t("labelSelectorResultsPage", { number: 100 }),
            value: 100
        },
        {
            label: t("labelSelectorResultsPage", { number: 200 }),
            value: 200
        },
    ]);

    const currentPage = defineModel<number>("currentPage");

    const pageSize = defineModel<number>("pageSize");
</script>

<template>
    <n-flex justify="space-between" class="doneo-pagination-container">
        <div class="doneo-pagination-total-results-container">
            <span>
                <slot name="total-results-label" :total-results="props.totalResults">
                    {{ t("Total results:") }} {{ props.totalResults }}
                </slot>
            </span>
        </div>
        <n-pagination v-model:page="currentPage" v-model:page-size="pageSize" :page-count="totalPages"
            :page-sizes="pageSizes" show-size-picker :page-slot="8">
            <template #prefix="{ page, pageCount }">
                {{ t("labelPageOfTotalPages", { currentPage: page, totalPages: pageCount }) }}
            </template>
        </n-pagination>
    </n-flex>
</template>

<style lang="css" scoped>
    .doneo-pagination-container {
        padding: 4px;
        background-color: rgba(250, 250, 252, 1);
        margin-bottom: 8px;
        border: 1px solid;
        border-color: rgb(239, 239, 245);
        border-radius: 3px;
    }

    .doneo-doneo-pagination-total-results-container {
        padding-top: 2px;
        padding-left: 2px;
    }
</style>