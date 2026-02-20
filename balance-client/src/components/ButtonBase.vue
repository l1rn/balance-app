<template>
    <div>
        <button>
            <img :src="props.icon" alt="" srcset="">
            {{ props.title }}
        </button>
    </div>
</template>

<script setup>

const props = defineProps({
    title: {
        type: String
    },
    icon: {
        type: String
    }
})
</script>

<style scoped>

div button {
    border: 0;
    outline: 0;
    cursor: pointer;
    border-radius: 8px;
    padding: 0.25rem;
    display: flex;
    align-items: center;
    justify-content: center;
    background: #3d3d3d;
    transition: all 0.3s ease;

}

div button:hover {
    background: #505050;
}

div button img {
    margin: 0;
    padding: 0;

}
</style>