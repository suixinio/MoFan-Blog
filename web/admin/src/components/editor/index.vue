<template>
  <div>
    <div id='vditor' class='vditor' />
  </div>
</template>

<script>

import Vditor from 'vditor'
import 'vditor/dist/index.css'

export default {

  props: {
    value: {
      type: String,
      default: ''
    },
    vdContent: {
      type: String,
      default: ''
    }

  },
  data () {
    return {
      contentEditor: '',
      content: ''
    }
  },
  created () {
    // this.cleanText()
    // this.setDefaultText()
  },
  mounted () {
    const that = this
    const options = {
      typewriterMode: true,
      after () {
        console.log('渲染完成')
      },
      ctrlEnter (md) {
        console.log('用户按下了 Ctrl+Enter，Markdown 内容为：\n' + md)
        that.content = md
      },
      input (md) {
        console.log('用户进行了输入，Markdown 内容为：\n' + md)
        that.content = md
      },
      blur (md) {
        console.log('用户离开了编辑器，Markdown 内容为：\n' + md)
        that.content = md
      },
      select (md) {
        console.log('用户选中了一段文字，内容为：\n' + md)
      },
      focus (md) {
        console.log('用户选中了编辑器，Markdown 内容为：\n' + md)
        that.content = md
      },
      esc (md) {
        console.log('用户按下了 ESC，Markdown 内容为：\n' + md)
        that.content = md
      },
      // width: this.isMobile ? '100%' : '80%',
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
  },
  methods: {
    // cleanText () {
    //   localStorage.setItem('vditorvditor', '')
    // },
    // saveText (value) {
    //   localStorage.setItem('vditorvditor', value)
    // }
  },
  watch: {
    vdContent (newV) {
      console.log('vdContent' + newV)
      // this.saveText(newV)
      this.content = newV
      // this.contentEditor.setValue(newV)
    },
    value (newV) {
      console.log('value:', newV)
      // this.setText(newV)
    },
    content (newV) {
      console.log('content:' + newV)
      this.$emit('input', newV)
    }
  }
}
</script>

<style scoped>
.vditor {
  /*position: absolute;*/
  /*top: @header-height;*/
  /*max-width: @max-body-width;*/
  /*width: 80%;*/
  /*height: calc(100vh - 100px);*/
  /*margin: 20px auto;*/
  /*text-align: left;*/
  height: 70vh ;
}
</style>
