import { PayloadAction, createSlice } from '@reduxjs/toolkit';
import { IUser } from '../../types/User';
// import { createUser, getStudentInfo, updateUser } from "../../thunks/user";
// import { ICompany, IUser } from "../../../types/Type";

export interface UserState {
  data: IUser;
  error?: string;
  isLoading?: boolean;
}

export const initialState: UserState = {
  data: {
    id: '000000000000000000000',
    token: '',
    created_at: new Date().toUTCString(),
    updated_at: new Date().toUTCString(),
    picture:
      'https://images.unsplash.com/photo-1508214751196-bcfd4ca60f91?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=255&q=80',
    given_name: '',
    family_name: '',
    email: '',
  },
  isLoading: false,
  error: '',
};

export const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {},
  });
export default userSlice.reducer;
