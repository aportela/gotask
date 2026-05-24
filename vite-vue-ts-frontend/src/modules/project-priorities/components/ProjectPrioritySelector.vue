<script setup lang="ts">

    import { ref, shallowRef, reactive, computed, onMounted, onBeforeUnmount } from 'vue';
    import { NSelect, type SelectOption, type SelectSize } from 'naive-ui';

    import { type AjaxStateInterface, defaultAjaxState, defaultAjaxStateRunning } from '../../../shared/types/ajaxState';
    import { projectPriorityService } from '../services/project-priority';
    import type { SearchRequest, ProjectPriorityResponse } from '../types/dto';
    import { Sort } from '../../../shared/types/models/sort';
    import { appBus } from '../../../shared/composables/bus';
    import { handleAPIError } from '../../../api/client/errorHandler';

    interface ProjectPrioritySelectorProps {
        placeholder?: string;
        size?: SelectSize;
    }

    const state: AjaxStateInterface = reactive({ ...defaultAjaxState });

    const isDisabled = computed(() => state.ajaxRunning);

    const projectPriorityId = defineModel<string | null>('id');

    const props = defineProps<ProjectPrioritySelectorProps>();

    const sort = ref<Sort>(new Sort("name", "ASC"));

    const options = shallowRef<SelectOption[]>([]);

    const onRefresh = async () => {
        Object.assign(state, defaultAjaxStateRunning);
        try {
            const payload: SearchRequest = {
                pager: {
                    currentPage: 1,
                    resultsPage: 0,
                },
                order: {
                    field: sort.value.field,
                    sort: sort.value.order,
                },
                filter: {
                    name: undefined,
                }
            };
            const response = await projectPriorityService.search(payload);
            options.value = response.projectPriorities.map((projectPriority: ProjectPriorityResponse) => ({ label: projectPriority.name, value: projectPriority.id }));
        } catch (error: unknown) {
            options.value.length = 0;
            state.ajaxErrors = true;
            handleAPIError(error,
                (apiError) => {
                    switch (apiError.response?.status) {
                        case 401:
                            state.ajaxErrors = false;
                            appBus.emit({ type: "reauthRequired", payload: { emitter: "ProjectPrioritySelector.onRefresh" } });
                            break;
                        default:
                            //state.ajaxErrorMessage = t("modules.projectPriority.components.ManageProjectPrioritiesPage.errors.refreshError");
                            break;
                    }
                },
                (fatalError) => {
                    //state.ajaxErrorMessage = t("modules.projectPriority.components.ManageProjectPrioritiesPage.errors.refreshError");
                    console.error("Unhandled API error", { file: "ProjectPrioritySelector.vue", method: "onRefresh" }, { err: fatalError });
                });
        }
        finally {
            state.ajaxRunning = false;
        }
    };

    let stopBusReauthListener: () => void;

    onMounted(() => {
        stopBusReauthListener = appBus.on("reauthValidNotify", async (payload) => {
            if (payload.to.includes("ProjectPrioritySelector.onRefresh")) {
                onRefresh();
            }
        });
        onRefresh();
    });

    onBeforeUnmount(() => {
        stopBusReauthListener();
    });

</script>

<template>
    <n-select v-model:value="projectPriorityId" :options="options" :placeholder="props.placeholder" :size="props.size"
        :disabled="isDisabled" />
</template>

<style lang="css" scoped></style>