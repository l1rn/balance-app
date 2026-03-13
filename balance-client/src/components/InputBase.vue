<template>
    <div class="input-base">
        <input 
        v-bind="$attrs"
        v-model="localValue" 
        :type="props.type"
        :readonly="props.type === 'list'"
        :placeholder="props.placeholder"
        :class="{'selector-mode': props.type === 'list'}">
        <template v-if="props.type === 'list' && props.items">
            <div class="list-container">
                <ul>
                    <li v-for="i in props.items">
                        {{ i }}
                    </li>
                </ul>
            </div>
        </template>
    </div>
</template>

<script setup lang="ts">

const props = defineProps({
  type: {
    type: String,
    default: "text",
    validator: (value) => ["text", "password", "list"].includes(value)
  },
  placeholder: {
    type: String
  },
  items: {
    type: Array
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
    border: 2px solid #535353;
}

.selector-mode {
    background: #535353;
    cursor: pointer;
    border: 2px solid #535353;
}

.list-container {
    
}
</style>