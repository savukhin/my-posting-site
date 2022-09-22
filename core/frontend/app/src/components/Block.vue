<script lang="ts">
export default {
  props: {
    height: {
        type: Number,
        require: false
    } 
  }
}
</script>

<template>
    <div class="block">
        <slot></slot>
    </div>
</template>
