<script setup lang="ts">
    import { reactive, watch, onMounted, type CSSProperties, shallowRef } from 'vue';
    import { useI18n } from "vue-i18n";
    import { type SelectOption, NSelect } from 'naive-ui';
    import { api } from '../../composables/api';
    import { type SearchWorkspacesResponse } from '../../types/apiResponses';
    import { type AjaxStateInterface, defaultAjaxState } from "../../types/ajaxState";
    import { useCurrentWorkspaceStore } from '../../stores/currentWorkspace';

    const { t } = useI18n();

    const currentWorkspaceStore = useCurrentWorkspaceStore();

    interface WorkspaceSelectorProps {
        style?: string | CSSProperties;
    };

    const props = defineProps<WorkspaceSelectorProps>();

    const emit = defineEmits(['error'])

    const selectedWorkspace = defineModel<string | null>('value');

    const options = shallowRef<SelectOption[]>([]);

    const state: AjaxStateInterface = reactive({ ...defaultAjaxState });

    watch(selectedWorkspace, (val) => {
        if (val) {
            currentWorkspaceStore.set(val);
        }
    });

    const findValidWorkspace = (options: SelectOption[], stored?: string | null) => {
        if (!options.length) return null;

        if (stored && options.some(o => o.value === stored)) {
            return stored;
        }

        return options[0].value as string;
    };

    const onRefresh = () => {
        Object.assign(state, defaultAjaxState);
        state.ajaxRunning = true;
        api.workspace.search().then((response: SearchWorkspacesResponse) => {
            options.value = response.data.workspaces.map((item) => ({
                value: item.id,
                label: item.name,
                disabled: false
            }));
            selectedWorkspace.value = findValidWorkspace(
                options.value,
                currentWorkspaceStore.currentWorkspaceId
            );
        }).catch((error: any) => {
            emit("error", { error: error });
        }).finally(() => {
            state.ajaxRunning = false;
        });
    };

    onMounted(() => {
        onRefresh();
    });
</script>

<template>
    <n-select :placeholder="t('Workspace')" v-model:value="selectedWorkspace" :options="options" :style="style"
        :disabled="state.ajaxRunning" :loading="state.ajaxRunning"></n-select>
</template>

<style lang="css" scoped></style>