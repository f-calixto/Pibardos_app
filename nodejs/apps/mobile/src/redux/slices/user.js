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
      return rejectWithValue(error.response.data.errors)
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
      return rejectWithValue(error.response.data.errors)
    }
  })

const userSlice = createSlice({
  name: 'user',
  initialState: {
    status: 'idle',
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
      state.errors = action.payload
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
      state.errors = action.payload
    }
  }
})

// export const { registerUser } = userSlice.actions
export default userSlice.reducer
