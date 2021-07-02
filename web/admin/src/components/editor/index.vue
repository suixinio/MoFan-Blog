<template>
  <div>
    <!--    <Editor :init='init' v-model='content'></Editor>-->

    <div id='vditor' class='vditor' />
  </div>
</template>

<script>
// import Editor from '@tinymce/tinymce-vue'
// eslint-disable-next-line no-unused-vars
// import tinymce from './tinymce.min'
// import './icons/default/icons.min'
// import './themes/silver/theme.min'
// import './langs/zh_CN'

import Vditor from 'vditor'
import 'vditor/dist/index.css'

export default {
  // components: { Editor },

  props: {
    value: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      contentEditor: ''
    }
  },
  created () {
    this.setDefaultText()
    console.log(this.value)
  },
  mounted () {
    const that = this
    const options = {
      typewriterMode: true,
      after () {
        console.log('渲染完成')
        that.contentEditor.setValue(that.value)
      },
      ctrlEnter (md) {
        console.log('用户按下了 Ctrl+Enter，Markdown 内容为：\n' + md)
        that.content = md
        that.$emit('input', md)
      },
      input (md) {
        console.log('用户进行了输入，Markdown 内容为：\n' + md)
        that.content = md
        that.$emit('input', md)
      },
      blur (md) {
        console.log('用户离开了编辑器，Markdown 内容为：\n' + md)
        that.content = md
        that.$emit('input', md)
      },
      select (md) {
        console.log('用户选中了一段文字，内容为：\n' + md)
        // this.content = md
      },
      focus (md) {
        console.log('用户选中了编辑器，Markdown 内容为：\n' + md)
        that.content = md
        that.$emit('input', md)
      },
      esc (md) {
        console.log('用户按下了 ESC，Markdown 内容为：\n' + md)
        that.content = md
        that.$emit('input', md)
      },
      width: this.isMobile ? '100%' : '80%',
      height: '0',
      tab: '\t',
      counter: '999999',
      mode: 'sv',
      preview: {
        delay: 100,
        show: !this.isMobile
      },
      outline: true
    }
    this.contentEditor = new Vditor('vditor', options)
    // this.contentEditor.setValue('123123')
    this.contentEditor.clearCache()
  },
  methods: {
    update (str) {
      this.contentEditor.setValue('123123')
    },
    setDefaultText () {
      const savedMdContent = localStorage.getItem('vditorvditor') || ''
      if (!savedMdContent.trim()) {
        localStorage.setItem('vditorvditor', '12312312313')
      }
    }
  },
  watch: {
    value (newV) {
      console.log('value' + newV)
      this.content = newV
      this.contentEditor.setValue(newV)
    },
    content (newV) {
      console.log('content:' + newV)
      this.$emit('input', newV)
    }
  }
}
</script>

<style scoped>
@import url('./skins/ui/oxide/skin.min.css');
</style>
