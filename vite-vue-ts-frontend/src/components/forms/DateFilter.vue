<script setup lang="ts">
    import { ref, watch } from 'vue';
    import { NSelect, NDatePicker } from 'naive-ui';


    const filterDateOptions = [
        { label: 'Any date', value: 0 },
        { label: 'Today', value: 1 },
        { label: 'Yesterday', value: 2 },
        { label: 'This week', value: 3 },
        { label: 'Custom date', value: 4 }
    ];


    const selectorValue = ref<number>(0)
    const datepickerValue = ref<number | null>(null)
    const showPicker = ref<boolean>(false);

    const onClearDate = () => {
        datepickerValue.value = null
        selectorValue.value = 0
        showPicker.value = false
    };

    watch(selectorValue, async (val) => {
        if (val !== 4) {
            datepickerValue.value = null
            showPicker.value = false
            return
        }
        showPicker.value = true
    });

    watch(showPicker, (newValue: boolean) => {
        if (!newValue && !datepickerValue.value) {
            selectorValue.value = 0;
        }
    });

</script>

<template>
    <n-select v-if="selectorValue !== 4" v-model:value="selectorValue" :options="filterDateOptions" size="small" />

    <n-date-picker v-else v-model:value="datepickerValue" type="date" clearable size="small" v-model:show="showPicker"
        @clear="onClearDate" :actions="['clear']" />
</template>

<style lang="css" scoped></style>