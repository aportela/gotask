<script setup lang="ts">
    const title = "Title";

</script>

<template>
    <n-table size="small">
        <caption class="table-caption">
            <n-grid :cols="2" align="center">
                <n-grid-item style="text-align: left;">
                    <span class="table-caption-title">{{ title }}</span>
                </n-grid-item>
                <!--
                <n-grid-item style="display: flex; justify-content: flex-end;">
                    <n-button @click="onAddProjectType" :disabled="state.ajaxRunning">
                        <template #icon>
                            <IconPlus />
                        </template>
{{ t("Add") }}
</n-button>
</n-grid-item>
-->
            </n-grid>
        </caption>

        <slot name="thead">
            <thead></thead>
        </slot>
        <slot name="tbody">
            <tbody></tbody>
        </slot>
    </n-table>
</template>

<style lang="css" scoped>
    .table-caption {
        padding-bottom: 4px;
    }

    .table-caption-title {
        font-weight: 700;
        font-size: large;

    }
</style>