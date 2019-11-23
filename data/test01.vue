<template>
  <!-- some html code -->
  <div> ... </div>
  <!-- some html code with v-command (like v-if, v-for) and injected js command -->
  <div v-if="showTag === true">
    <div v-for="eachItem in data"> ... </div>
  </div>
</template>

<script>
// local functions and variables
var a = 1, b = 2
function test () {}
var d = function () {}
var e = () => {}

export default {
  name: 'name of a component',
   props: {
     'data': {
       'type': [Object],
       'default': function () {
         return {
           'columns': [],
           'fixColumnCount': null
         }
       }
     }
   },
  data () {
    return {
      'show': false
    }
  },
  computed: {
    style () {
    },
  },
  'methods': { // this is usually a LARGE dict and will take 50% of the total file
    someMethod0 () {
    },
    someMethod1 () {
    },
    someMethod0 () {
    }
  },
  created () {
  },
  components: {MTableRow}
}
</script>

<style>
</style>
