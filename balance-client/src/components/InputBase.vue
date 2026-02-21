<template>
    <div class="input-wrapper">
        <input 
        v-model="localValue" 
        :type="props.type"
        :placeholder="props.placeholder">
    </div>
</template>

<script setup lang="ts">

const props = defineProps({
  type: {
    type: String,
    default: "text",
    validator: (value) => ["text", "password"].includes(value)
  },
  placeholder: {
    type: String
  }
})

const localValue = defineModel({
    default: ''
})
</script>

<style>

input {
    width: 100%;
    outline: 0;
    border: 2px solid #1f1f1f;
    text-indent: 0.5rem;
    box-sizing: border-box;
    border-radius: 8px;
    transition: all 0.2s;
}

input:focus {
    border: 2px solid #323036;
}

</style>