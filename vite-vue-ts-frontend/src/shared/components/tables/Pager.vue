<script setup lang="ts">
    import { NFlex, NPagination } from 'naive-ui';

    interface PagerProps {
        totalResults: number;
        totalPages: number;
    };

    const props = defineProps<PagerProps>();

    const pageSizes = [
        {
            label: 'All results',
            value: 0
        },
        {
            label: '5 results/page',
            value: 5
        },
        {
            label: '10 results/page',
            value: 10
        },
        {
            label: '20 results/page',
            value: 20
        },
        {
            label: '50 results/page',
            value: 50
        },
        {
            label: '100 results/page',
            value: 100
        },
        {
            label: '200 results/page',
            value: 200
        },
    ];

    const currentPage = defineModel<number>("currentPage");

    const pageSize = defineModel<number>("pageSize");

</script>

<template>
    <n-flex justify="space-between" class="pagination-container">
        <div style="padding-top: 2px; padding-left: 2px;">
            <span>Total results: {{ props.totalResults }}</span>
        </div>
        <n-pagination v-model:page="currentPage" :page-count="totalPages" v-model:page-size="pageSize" show-size-picker
            :page-sizes="pageSizes" :page-slot="8">
            <template #prefix="{ page, pageCount }">
                Page {{ page }} of {{ pageCount }}
            </template>
        </n-pagination>
    </n-flex>
</template>

<style lang="css" scoped>
    .pagination-container {
        padding: 4px;
        background-color: rgba(250, 250, 252, 1);
        margin-bottom: 8px;
        border: 1px solid;
        border-color: rgb(239, 239, 245);
        border-radius: 3px;
    }
</style>