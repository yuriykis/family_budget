<template>
    <v-toolbar style="border-radius: 80px;">
        <v-text-field hide-details prepend-icon="mdi-magnify" single-line color="#363636"></v-text-field>
        <v-btn icon>
            <v-icon>mdi-dots-vertical</v-icon>
        </v-btn>
    </v-toolbar>
</template>

<script>
export default {
    name: 'SearchBar',
    data: () => ({
    }),
}
</script>
