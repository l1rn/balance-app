<template>
    <div class="input-base">
        <input 
            v-bind="$attrs"
            v-model="localValue" 
            :type="props.type === 'list' ? 'text' : props.type"
            :readonly="props.type === 'list'"
            @click="isOpen = !isOpen"
            @blur="isOpen = false" 
            :placeholder="props.placeholder"
            :class="{'selector-mode': props.type === 'list'}"
        >
        
        <div v-if="props.type === 'list' && props.items && isOpen" class="list-container">
            <ul>
                <li v-for="i in props.items" :key="i" @mousedown="selectItem(i)">
                    {{ i }}
                </li>
            </ul>
        </div>
    </div>
</template>
<script setup>
import { ref } from 'vue';

const isOpen = ref(false);

const selectItem = (item) => {
    localValue.value = item; 
    isOpen.value = false;   
};

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
.input-base {
    position: relative;
    width: 100%;
}
.list-container {
    position: absolute;
    top: 105%; 
    left: 0;
    width: 100%;
    background: #323036;
    border: 2px solid #535353;
    box-sizing: border-box;
    border-radius: 8px;
    z-index: 999; 
    overflow: hidden;
}

.list-container ul {
    margin: 0;
    padding: 0;
    list-style: none;
}

.list-container ul li {
    padding: 0.5rem 1rem;
    cursor: pointer;
    transition: background 0.2s;
    text-align: left;
    color: white;
}

.list-container ul li:hover {
    background: #535353; 
}
</style>