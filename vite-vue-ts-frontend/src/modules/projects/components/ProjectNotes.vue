<script setup lang="ts">
    import { shallowRef, reactive, onMounted, onBeforeUnmount, watch, type CSSProperties } from "vue";
    import { useI18n } from "vue-i18n";

    import { NSpace, NCard } from "naive-ui";

    import { type AjaxStateInterface, defaultAjaxState, defaultAjaxStateRunning } from '../../../shared/types/ajaxState';
    import { useLoadingStore } from '../../../stores/loading';
    //    import { useNotify } from '../../../shared/composables/notification';
    import { appBus } from '../../../shared/composables/bus';
    import { Note } from "../../notes/models/note.ts";
    import { noteService } from "../../notes/services/note.ts";
    import { handleAPIError } from '../../../api/client/errorHandler';
    import type { SearchResponse } from "../../notes/types/dto.ts";
    import AvatarUserName from "../../../shared/components/AvatarUserName.vue";


    interface ProjectNotesProps {
        style?: string | CSSProperties;
        projectId: string | null;
    }

    const props = defineProps<ProjectNotesProps>();

    const { t } = useI18n();
    //const { notify } = useNotify();
    const loadingStore = useLoadingStore();

    const state: AjaxStateInterface = reactive({ ...defaultAjaxState });

    const items = shallowRef<Note[]>([]);

    const itemCount = defineModel<number>("itemCount", { default: 0 });

    watch(state, (newValue: AjaxStateInterface) => {
        loadingStore.set(newValue.ajaxRunning);
    });

    watch(() => props.projectId, (newValue, oldValue) => {
        if (!oldValue && newValue) {
            onRefresh();
        }
    });

    const onRefresh = async () => {
        if (props.projectId) {
            Object.assign(state, defaultAjaxStateRunning);
            try {
                const results: SearchResponse = await noteService.search(props.projectId);
                items.value = results.notes.map((note) => new Note(note));
                itemCount.value = items.value?.length ?? 0;
            } catch (error: unknown) {
                state.ajaxErrors = true;
                handleAPIError(error,
                    (apiError) => {
                        switch (apiError.response?.status) {
                            case 401:
                                state.ajaxErrors = false;
                                appBus.emit({ type: "reauthRequired", payload: { emitter: "ProjectNotes.onRefresh" } });
                                break;
                            default:
                                state.ajaxErrorMessage = t("modules.projectPermission.components.projectPermissions.errors.refreshError");
                                break;
                        }
                    },
                    (fatalError) => {
                        state.ajaxErrorMessage = t("modules.projectPermission.components.projectPermissions.errors.refreshError");
                        console.error("Unhandled API error", { file: "ProjectNotes.vue", method: "onRefresh" }, { err: fatalError });
                    });
            } finally {
                state.ajaxRunning = false;
            }
        } else {
            console.error("project id not set", { file: "ProjectNotes.vue", method: "onRefresh" });
        }
    };

    let stopBusReauthListener: () => void;

    onMounted(() => {
        if (props.projectId) {
            onRefresh();
        }
        stopBusReauthListener = appBus.on("reauthValidNotify", async (payload) => {
            if (payload.to.includes("ProjectNotes.onRefresh")) {
                onRefresh();
            } else if (payload.to.includes("ProjectNotes.onDelete")) {
                //onDelete(selectedItem.value);
            }
        });
    });

    onBeforeUnmount(() => {
        stopBusReauthListener();
    });
</script>

<template>
    <n-card bordered :style="props.style">
        <n-space vertical size="large" style="margin-right: 12px;">
            <n-card v-for="note, index in items" :key="note.id ?? index" size="small" bordered>
                <div class="note-header">
                    <div class="note-user">
                        <AvatarUserName :user-id="note.user.id" :user-name="note.user.name" />
                    </div>
                    <span class="note-date">
                        {{ note.createdAt?.toLocaleString() }}
                    </span>
                </div>
                <div class="note-content">
                    <p>
                        random note text {{ index }}
                    </p>
                    {{ note.body }}
                </div>
            </n-card>
        </n-space>
    </n-card>
</template>

<style lang="css" scoped>
    .note-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 12px;
    }

    .note-user {
        display: flex;
        align-items: center;
        gap: 8px;
        font-weight: 500;
    }

    .note-date {
        font-size: 12px;
        color: #999;
    }

    .note-content {
        font-size: 14px;
        white-space: pre-line;
    }


</style>