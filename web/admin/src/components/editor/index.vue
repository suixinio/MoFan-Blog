<template>
  <div>
    <Editor :init='init' v-model='content'></Editor>
  </div>
</template>

<script>
import Editor from '@tinymce/tinymce-vue'
// eslint-disable-next-line no-unused-vars
import tinymce from './tinymce.min'
import './icons/default/icons.min'
import './themes/silver/theme.min'
import './langs/zh_CN'

export default {
  components: { Editor },
  props: {
    value: {
      type: String,
      default: ''
    }
  },

  data () {
    return {
      init: {
        language: 'zh_CN'
      },
      content: this.value
    }
  },
  watch: {
    value (newV) {
      this.content = newV
    },
    content (newV) {
      this.$emit('input', newV)
    }

  }
}
</script>

<style scoped>
@import url('./skins/ui/oxide/skin.min.css');
</style>
