<script setup lang="ts">
    import { NMenu, NIcon, type MenuOption } from 'naive-ui'
    import { IconUserCheck, IconBug, IconSitemap, IconFileAnalytics, IconSettings, IconUserCircle, IconPresentation, IconChartHistogram, IconUsers, IconBookmark, IconFlagBolt, IconAdjustmentsBolt, IconLogout, IconId } from '@tabler/icons-vue';
    import { ref, watch, h } from 'vue'
    import type { Component } from 'vue'
    import { useRouter } from "vue-router";
    import { RouterLink } from 'vue-router'


    const router = useRouter();
    const currentTab = ref<string | number>("home");

    const commonIconSize = 20;

    watch(() => currentTab.value, (newValue) => {
        router.push(
            { name: newValue.toString() }
        ).catch((e) => {
            console.error(e);
        });
    });

    function renderIcon(icon: Component) {
        return (size = commonIconSize) =>
            () =>
                h(NIcon, { size }, {
                    default: () => h(icon)
                })
    }

    function handleUpdateValue(_key: string, _item: MenuOption) {
    }

    const menuOptions: MenuOption[] = [
        {
            label: () =>
                h(
                    RouterLink,
                    {
                        to: {
                            name: 'home',
                            params: {
                            }
                        }
                    },
                    { default: () => 'Overview' }
                ),
            key: 'home',
            icon: renderIcon(IconPresentation)(commonIconSize),
        },
        {
            label: () =>
                h(
                    RouterLink,
                    {
                        to: {
                            name: 'projects',
                            params: {
                            }
                        }
                    },
                    { default: () => 'Projects' }
                ),
            key: 'projects',
            icon: renderIcon(IconSitemap)(commonIconSize),
        },
        {
            label: "Tasks",
            key: 'tasks',
            disabled: true,
            icon: renderIcon(IconBug)(commonIconSize)
        },
        {
            label: "Reports",
            key: "reports",
            disabled: true,
            icon: renderIcon(IconFileAnalytics)(commonIconSize)
        },
        {
            label: "Charts",
            key: 'charts',
            disabled: true,
            icon: renderIcon(IconChartHistogram)(commonIconSize)
        },
        {
            label: "Settings",
            key: 'settings',
            show: false,
            icon: renderIcon(IconSettings)(commonIconSize),
            children: [
                {
                    label: () =>
                        h(
                            RouterLink,
                            {
                                to: {
                                    name: 'users',
                                    params: {
                                    }
                                }
                            },
                            { default: () => 'Users' }
                        ),
                    key: 'users',
                    icon: renderIcon(IconUsers)(commonIconSize)
                },
                {
                    label: 'Roles',
                    key: 'roles',
                    disabled: true,
                    icon: renderIcon(IconUserCheck)(commonIconSize),
                },
                {
                    label: "Project",
                    key: 'projectSettings',
                    icon: renderIcon(IconSettings)(commonIconSize),
                    children: [
                        {
                            label: () =>
                                h(
                                    RouterLink,
                                    {
                                        to: {
                                            name: 'manageProjectTypes',
                                            params: {
                                            }
                                        }
                                    },
                                    { default: () => 'Type' }
                                ),
                            key: "projectTypes",
                            icon: renderIcon(IconBookmark)(commonIconSize)
                        },
                        {
                            label: () =>
                                h(
                                    RouterLink,
                                    {
                                        to: {
                                            name: 'manageProjectPriorities',
                                            params: {
                                            }
                                        }
                                    },
                                    { default: () => 'Priority' }
                                ),
                            key: "projectPriorities",
                            icon: renderIcon(IconFlagBolt)(commonIconSize)
                        },
                        {
                            label: () =>
                                h(
                                    RouterLink,
                                    {
                                        to: {
                                            name: 'manageProjectStatuses',
                                            params: {
                                            }
                                        }
                                    },
                                    { default: () => 'Status' }
                                ),
                            key: "projectStatuses",
                            icon: renderIcon(IconAdjustmentsBolt)(commonIconSize)

                        }
                    ]
                },
            ]
        },
        {
            label: "John Doe",
            key: 'myuser',
            icon: renderIcon(IconUserCircle)(commonIconSize),
            children: [
                {
                    label: "Profile",
                    key: "profile",
                    disabled: true,
                    icon: renderIcon(IconId)(commonIconSize),
                },
                {
                    label: "Logout",
                    key: "signout",
                    icon: renderIcon(IconLogout)(commonIconSize),
                }
            ]
        },
    ];
</script>

<template>
    <div class="header">
        <n-menu :collapsed-width="64" :collapsed-icon-size="commonIconSize" :options="menuOptions" mode="horizontal"
            @update:value="handleUpdateValue" style="width: max-content;" />
    </div>
</template>

<style lang="css" scoped>
    .header {
        height: 64px;
        display: flex;
        justify-content: center;
        align-items: center;
        padding: 0 10px;
        box-sizing: border-box;
        width: 100%;
    }

</style>