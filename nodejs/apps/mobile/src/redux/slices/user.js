import { createSlice, createAsyncThunk } from '@reduxjs/toolkit'
import { userService } from '../../services/user.service'

export const registerUser = createAsyncThunk('user/registerUser',
  async ({
    email,
    password,
    username,
    birthdate,
    country
  }, { rejectWithValue, dispatch }) => {
    try {
      await userService.registerUser({ email, password, username, birthdate, country })
      dispatch(loginUser({
        email,
        password
      }))
    } catch (error) {
      return rejectWithValue({
        statusCode: error.response.status,
        errors: error.response.data.errors
      })
    }
  })

export const loginUser = createAsyncThunk('user/loginUser',
  async ({
    email,
    password
  }, { rejectWithValue }) => {
    try {
      const response = await userService.loginUser({ email, password })
      return response.data
    } catch (error) {
      return rejectWithValue({
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
    accessToken: null,
    user: {
      id: null,
      email: null,
      username: null,
      country: null
    }
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

      if (action.payload.statusCode !== 400) {
        state.error = 'An unknown error occured'
      } else {
        state.errors = action.payload.errors
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
      state.accessToken = action.payload.accessToken
    },
    [loginUser.rejected]: (state, action) => {
      state.status = 'failed'

      if (action.payload.statusCode !== 400) {
        state.error = 'An unknown error occured'
      } else {
        state.errors = action.payload.errors
      }
    }
  }
})

// export const { registerUser } = userSlice.actions
export default userSlice.reducer
