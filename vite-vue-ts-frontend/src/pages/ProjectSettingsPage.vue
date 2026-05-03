<script setup lang="ts">
    import { NCard, NTabs, NTabPane } from 'naive-ui'
    import { default as ManageProjectTypesCard } from '../components/cards/settings/ManageProjectTypesCard.vue';
    import { default as ManageProjectStatusesCard } from '../components/cards/settings/ManageProjectStatusesCard.vue';
    import { default as ManageProjectPrioritiesCard } from '../components/cards/settings/ManageProjectPrioritiesCard.vue';
</script>

<template>
    <n-card style="margin: 1em;">
        <n-tabs placement="left" type="line" :bar-width="48">
            <n-tab-pane name="Project types" tab="Project types">
                <ManageProjectTypesCard />
            </n-tab-pane>
            <n-tab-pane name="Project priorities" tab="Project priorities">
                <ManageProjectPrioritiesCard />
            </n-tab-pane>
            <n-tab-pane name="Project statuses" tab="Project statuses">
                <ManageProjectStatusesCard />
            </n-tab-pane>
        </n-tabs>
    </n-card>
</template>

<style lang="css" scoped></style>