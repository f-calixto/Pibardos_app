import { createSlice, createAsyncThunk } from '@reduxjs/toolkit'
import { authService } from '@Services/auth.service'
import { secureStoreService } from '@Services/secureStore.service'
import { usersService } from '@Services/users.service'

export const registerUser = createAsyncThunk('user/registerUser',
  async ({ email, password, username, birthdate, country }, thunkAPI) => {
    try {
      await authService.registerUser({ email, password, username, birthdate, country })
      thunkAPI.dispatch(loginUser({
        email,
        password
      }))
    } catch (error) {
      return thunkAPI.rejectWithValue({
        error: error.response.data.error,
        errors: error.response.data.errors
      })
    }
  })

export const loginUser = createAsyncThunk('user/loginUser',
  async ({ email, password }, thunkAPI) => {
    try {
      const response = await authService.loginUser({ email, password })

      // save user data into secure store for persist user's session
      await secureStoreService.saveUser({
        id: response.data.id,
        accessToken: response.data.accessToken
      })

      const { country, avatar } = await usersService.getUser(response.data.id).data

      return { ...response.data, country, avatar }
    } catch (error) {
      return thunkAPI.rejectWithValue({
        statusCode: error.response.status,
        errors: error.response.data.errors
      })
    }
  })

export const setUser = createAsyncThunk('user/setUser',
  async ({ id, accessToken }, thunkAPI) => {
    try {
      const response = await usersService.getUser(id)
      return { ...response.data, accessToken }
    } catch (error) {
      return thunkAPI.rejectWithValue({
        statusCode: error.response.status,
        errors: error.response.data.errors
      })
    }
  })

const userSlice = createSlice({
  name: 'user',
  initialState: {
    status: 'idle',
    error: null,
    errors: null,
    isLoggedIn: false,
    loggedUser: {
      accessToken: null,
      id: null,
      email: null,
      username: null,
      country: null,
      avatar: null
    }
  },
  reducers: {
  },
  extraReducers: {
    /* registerUser reducers */
    [registerUser.pending]: (state, action) => {
      state.status = 'loading'
      state.errors = null
    },
    [registerUser.fulfilled]: (state, action) => {
      state.status = 'succeeded'
      state.errors = null
    },
    [registerUser.rejected]: (state, action) => {
      state.status = 'failed'
      state.errors = action.payload.errors || null

      // TODO: refactor this block to avoid repeated code
      // if server not return an errors array, then set an error message
      // returned from server or set a generic error message.
      if (!action.payload.errors || action.payload.errors.length === 0) {
        state.error = action.payload.error || 'An unknown error occured'
      }
    },

    /* loginUser reducers */
    [loginUser.pending]: (state, action) => {
      state.status = 'loading'
      state.errors = null
    },
    [loginUser.fulfilled]: (state, action) => {
      state.status = 'succeeded'
      state.errors = null
      state.isLoggedIn = true
      state.loggedUser.accessToken = action.payload.accessToken

      state.loggedUser.id = action.payload.id
      state.loggedUser.email = action.payload.email
      state.loggedUser.username = action.payload.username
    },
    [loginUser.rejected]: (state, action) => {
      state.status = 'failed'
      state.errors = action.payload.errors || null

      // TODO: refactor this block to avoid repeated code
      // if server not return an errors array, then set an error message
      // returned from server or set a generic error message.
      if (!action.payload.errors || action.payload.errors.length === 0) {
        state.error = action.payload.error || 'An unknown error occured'
      }
    },

    /* setUser reducers */
    [setUser.pending]: (state, action) => {
      state.status = 'loading'
      state.errors = null
    },
    [setUser.fulfilled]: (state, action) => {
      state.status = 'succeeded'
      state.errors = null
      state.isLoggedIn = true
      state.loggedUser.accessToken = action.payload.accessToken

      state.loggedUser.id = action.payload.id
      state.loggedUser.email = action.payload.email
      state.loggedUser.username = action.payload.username
      state.loggedUser.country = action.payload.country
      state.loggedUser.avatar = action.payload.avatar
    },
    [setUser.rejected]: (state, action) => {
      state.status = 'failed'
      state.errors = action.payload.errors || null

      // TODO: refactor this block to avoid repeated code
      // if server not return an errors array, then set an error message
      // returned from server or set a generic error message.
      if (!action.payload.errors || action.payload.errors.length === 0) {
        state.error = action.payload.error || 'An unknown error occured'
      }
    }
  }
})

// export const { } = userSlice.actions
export default userSlice.reducer
