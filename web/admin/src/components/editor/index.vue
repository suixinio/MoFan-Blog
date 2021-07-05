<template>
  <div>
    <div id='vditor' class='vditor' />
  </div>
</template>

<script>

import Vditor from 'vditor'
import 'vditor/dist/index.css'
import { Url } from '../../plugin/http'

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
    this.headers = {
      Authorization: `Bearer ${window.sessionStorage.getItem('token')}`
    }
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
      cache: {
        // enable: false
      },
      upload: {
        accept: 'image/*,.mp3, .wav, .rar',
        headers: { Authorization: `Bearer ${window.sessionStorage.getItem('token')}` },
        url: Url + 'upload',
        fieldName: 'file',
        multiple: false,
        // linkToImgUrl: Url + 'upload',
        // linkToImgCallback: (responseText) => {
        //   console.log('rerere', responseText)
        // },
        success: (editor, msg) => {
          var json = JSON.parse(msg)
          console.log(editor, json.url)
          that.contentEditor.insertValue(`![](${json.url})`)
        },
        format: (files, responseText) => {
          console.log(files, responseText)
        }
        // linkToImgUrl: '/api/upload/fetch',
        // filename (name) {
        //   return name.replace(/[^(a-zA-Z0-9\u4e00-\u9fa5\.)]/g, '').replace(/[\?\\/:|<>\*\[\]\(\)\$%\{\}@~]/g, '').replace('/\\s/g', '')
        // }
      },
      // upload: {
      //   max: 5 * 1024 * 1024,
      //   // linkToImgUrl: 'https://sm.ms/api/upload',
      //   handler (file) {
      //     const formData = new FormData()
      //     for (const i in file) {
      //       formData.append('smfile', file[i])
      //     }
      //     const request = new XMLHttpRequest()
      //     request.open('POST', 'https://sm.ms/api/upload')
      //     request.onload = that.onloadCallback
      //     request.send(formData)
      //   }
      // },
      // upload: {
      //   url: that.Url + '/upload',
      //   accept: 'image/*,.mp3, .wav, .rar',
      //   token: that.base,
      //   linkToImgUrl: '/api/upload/fetch',
      //   filename (name) {
      //     // return name.replace(/[^(a-zA-Z0-9\u4e00-\u9fa5\.)]/g, '').replace(/[\?\\/:|<>\*\[\]\(\)\$%\{\}@~]/g, '').replace('/\\s/g', '')
      //   }
      // },

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
  height: 70vh;
}
</style>
