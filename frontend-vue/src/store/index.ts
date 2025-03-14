import { createStore } from 'vuex'

interface CartState {
  cartNum: number
}

export default createStore({
  state: (): CartState => ({
    cartNum: 0
  }),
  mutations: {
    SET_CART_NUM(state, num: number) {
      state.cartNum = num
    },
    INCREMENT_CART_NUM(state, delta: number = 1) {
      state.cartNum += delta
    },
    CLEAR_CART(state) {
      state.cartNum = 0
    }
  },
  actions: {
    initializeCart({ commit }) {
      const savedNum = localStorage.getItem('cartNum')
      if (savedNum) {
        commit('SET_CART_NUM', parseInt(savedNum))
      }
    },
    updateCartNum({ commit }, num: number) {
      commit('SET_CART_NUM', num)
      localStorage.setItem('cartNum', num.toString())
    }
  },
  getters: {
    cartNum: (state) => state.cartNum
  }
})