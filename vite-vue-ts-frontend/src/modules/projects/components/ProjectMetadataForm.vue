<script setup lang="ts">
    import { ref, type CSSProperties, nextTick } from 'vue';
    import { useI18n } from "vue-i18n";

    import { NCard, NForm, NFormItem, NInput, NInputGroup, NButton, NIcon, type InputInst } from 'naive-ui';

    import type { FormMode } from '../../../shared/types/form-mode';
    import { Project, MAX_KEY_LENGTH, MAX_SUMMARY_LENGTH } from "../models/project";
    import ProjectPrioritySelector from "../../project-priorities/components/ProjectPrioritySelector.vue";
    import ProjectStatusSelector from "../../project-statuses/components/ProjectStatusSelector.vue";
    import ProjectTypeSelector from "../../project-types/components/ProjectTypeSelector.vue";
    import AvatarUserName from '../../../shared/components/AvatarUserName.vue';
    import { IconCheck, IconDeviceFloppy, IconPencil } from '@tabler/icons-vue';

    interface ProjectFormProps {
        mode: FormMode;
        style?: string | CSSProperties;
        disabled?: boolean;
    }

    const props = defineProps<ProjectFormProps>();

    const emit = defineEmits(["save"]);


    const project = defineModel<Project>("project", { required: true });

    const { t } = useI18n();


    const keyEditMode = ref<boolean>(false);

    const summaryEditMode = ref<boolean>(false);

    const onSave = () => {
        emit("save");
    };

    const keyRef = ref<InputInst | null>(null);
    const summaryRef = ref<InputInst | null>(null);

    const onToggleSummaryMode = () => {
        summaryEditMode.value = !summaryEditMode.value;
        if (summaryEditMode) {
            nextTick(() => {
                summaryRef.value?.focus();
            });

        }
    };

    const onToggleKeyMode = () => {
        keyEditMode.value = !keyEditMode.value;
        if (keyEditMode) {
            nextTick(() => {
                keyRef.value?.focus();
            });

        }
    };

</script>

<template>
    <n-card bordered :style="props.style">
        <n-form-item label="Created by">
            <div class="note-user">
                <AvatarUserName :user-id="project.createdBy.id" :user-name="project.createdBy.name" />
            </div>
        </n-form-item>
        <n-form-item label="Created at">
            {{ project.createdAt.toLocaleString() }}
        </n-form-item>
        <n-form>
            <n-form-item label="Key">
                <n-input-group>
                    <n-input v-if="keyEditMode" v-model:value="project.key" :show-count="true"
                        :maxlength="MAX_KEY_LENGTH" ref="keyRef" :disabled="props.disabled" />
                    <div class="doneo-form-item-view" v-else>{{
                        project.summary
                    }}</div>
                    <n-button v-if="!keyEditMode" @click="onToggleKeyMode" :disabled="props.disabled">
                        <template #icon>
                            <n-icon :component="IconPencil" />
                        </template>
                    </n-button>
                    <n-button v-if="keyEditMode" @click="onToggleKeyMode" :disabled="props.disabled">
                        <template #icon>
                            <n-icon :component="IconCheck" />
                        </template>
                    </n-button>
                </n-input-group>
            </n-form-item>
            <n-form-item label="Summary">
                <n-input-group>
                    <n-input v-if="summaryEditMode" v-model:value="project.summary" :show-count="true"
                        :maxlength="MAX_SUMMARY_LENGTH" ref="summaryRef" :disabled="props.disabled" />
                    <div class="doneo-form-item-view" v-else>{{
                        project.summary
                    }}</div>
                    <n-button v-if="!summaryEditMode" @click="onToggleSummaryMode" :disabled="props.disabled">
                        <template #icon>
                            <n-icon :component="IconPencil" />
                        </template>
                    </n-button>
                    <n-button v-if="summaryEditMode" @click="onToggleSummaryMode" :disabled="props.disabled">
                        <template #icon>
                            <n-icon :component="IconCheck" />
                        </template>
                    </n-button>
                </n-input-group>
            </n-form-item>
            <n-form-item label="Description">
                <n-input v-model:value="project.description" type="textarea" clearable :disabled="props.disabled" />
            </n-form-item>
            <n-form-item label="Type">
                <ProjectTypeSelector v-model:id="project.type.id" :disabled="props.disabled" />
            </n-form-item>
            <n-form-item label="Priority">
                <ProjectPrioritySelector v-model:id="project.priority.id" :disabled="props.disabled" />
            </n-form-item>
            <n-form-item label="Status">
                <ProjectStatusSelector v-model:id="project.status.id" :disabled="props.disabled" />
            </n-form-item>
        </n-form>
        <n-button @click="onSave" :disabled="props.disabled">
            <template #icon>
                <n-icon :component="IconDeviceFloppy" color="red"></n-icon>
            </template>
            {{ t("shared.buttons.Save.label") }}
        </n-button>
    </n-card>
</template>

<style lang="css" scoped>

    .doneo-form-item-view {
        display: block;
        width: 100%;
        border: 1px solid #ccc;
        word-break: break-word;
        color: var(--n-text-color);
        border-radius: var(--n-border-radius);
        padding: 4px 12px;
        user-select: none;
    }
</style>