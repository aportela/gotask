<script setup lang="ts">
    import { NCard, NInput, NFlex, NButton, NColorPicker, NTag, NForm, NFormItem } from 'naive-ui';
    import { IconCancel, IconDeviceFloppy, IconTrash } from '@tabler/icons-vue';
    import { computed, onMounted } from 'vue';
    import { v7 as uuidv7 } from 'uuid';
    import type { EntityAction } from '../../types/common';

    import { ref } from 'vue';
    import type { CSSProperties } from 'vue';
    import { api } from '../../composables/api';
    import { type ProjectTypeInterface, ProjectTypeClass, maxNameLength } from '../../types/models/projectType';
    import { generateRandomSoftHexColor, getNaiveUITagColorProperty } from '../../composables/color';

    const emit = defineEmits(['add', 'update', 'delete', 'cancel'])

    interface ProjectTypeFormProps {
        mode: EntityAction;
        projectTypeId?: string;
        style?: string | CSSProperties;
    }

    const props = defineProps<ProjectTypeFormProps>();

    const saveButtonDisabled = computed<boolean>(() => {
        const inactiveMode = !(addMode.value || updateMode.value);
        const noName = !projectType.value.name;

        return inactiveMode || noName;
    });

    const addMode = computed<boolean>(() => props.mode === "add");
    const updateMode = computed<boolean>(() => props.mode === "update");
    const deleteMode = computed<boolean>(() => props.mode === "delete");

    const title = computed<string>(() => {
        switch (props.mode) {
            case "add":
                return "Add new project type";
            case "update":
                return "Update project type";
            case "delete":
                return "Delete project type";
            default:
                return "";
        }
    });

    const onSave = () => {
        if (addMode) {
            onAdd();
        } else if (updateMode) {
            onUpdate()
        } else {
            console.error("TODO");
        }
    }

    const onAdd = () => {
        emit('add')
    };

    const onUpdate = () => {
        emit('update')
    };

    const onDelete = () => {
        emit('delete')
    };

    const onCancel = () => {
        emit('cancel')
    }

    /*
const onAddProjectType = () => {
projectTypes.value.push(
    new ProjectType({ id: uuidv7(), name: "" })
);
nextTick(() => {
    if (tableFooter.value) {
        tableFooter.value?.scrollIntoView({
            behavior: 'smooth',
            block: 'end'
        });
    }
});
};
*/

    const projectType = ref<ProjectTypeInterface>(
        new ProjectTypeClass(
            { id: uuidv7(), name: "", hexColor: generateRandomSoftHexColor() }
        )
    );

    const getProjectType = (id: string) => {
        api.projectTypes.get(id).then((_response: any) => {
            projectType.value = _response.data.projectType;
        }).catch((_error: any) => { console.error(_error); }).finally(() => { });
    };

    onMounted(() => {
        if (props.mode == "update" || props.mode == "delete") {
            if (props.projectTypeId) {
                getProjectType(props.projectTypeId);
            } else {
                console.error("TODO")
            }
        }
    });

</script>

<template>
    <n-card :title="title" :style="style">
        <n-form>
            <n-form-item label="Name">
                <n-input placeholder="Project type name" v-model:value="projectType.name" :disabled="deleteMode"
                    :maxlength="maxNameLength" show-count clearable required autofocus></n-input>
            </n-form-item>
            <n-form-item label="Preview">
                <n-tag :color="getNaiveUITagColorProperty(projectType.hexColor)">
                    {{ projectType.name }}
                </n-tag>
            </n-form-item>
            <n-form-item label="Color">
                <n-color-picker :modes="['hex']" v-model:value="projectType.hexColor" />
            </n-form-item>
        </n-form>
        <template #action>
            <n-flex>
                <n-button @click="onSave" v-if="addMode || updateMode" :disabled="saveButtonDisabled">
                    <template #icon>
                        <IconDeviceFloppy />
                    </template>
                    Save
                </n-button>
                <n-button @click="onDelete" v-else-if="deleteMode">
                    <template #icon>
                        <IconTrash />
                    </template>
                    Delete
                </n-button>
                <n-button @click="onCancel">
                    <template #icon>
                        <IconCancel />
                    </template>
                    Cancel
                </n-button>
            </n-flex>
        </template>
    </n-card>
</template>

<style lang="css" scoped>
    .full-width-tag {
        display: block;
        width: 100%;
    }
</style>