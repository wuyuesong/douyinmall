// shims-vuex.d.ts
import { ComponentCustomProperties } from 'vue'
import { Store } from 'vuex'

declare module '@vue/runtime-core' {
  interface State {
    cartNum: number
  }

  interface ComponentCustomProperties {
    $store: Store<State>
  }
}