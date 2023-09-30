import { createSlice } from '@reduxjs/toolkit';
import { IAttendance } from '../../types/Attendance';
import { IUser } from '../../types/User';

export interface AttendanceState {
  data: IAttendance;
  error?: string;
  isLoading?: boolean;
}

export const initialState: AttendanceState = {
  data: {
    id: '000000000000000000000',
    author: {
      id: '000000000000000000000',
      token: '',
      created_at: new Date().toUTCString(),
      updated_at: new Date().toUTCString(),
      avatar: 'https://images.unsplash.com/photo-1508214751196-bcfd4ca60f91?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=255&q=80',
      given_name: '',
      family_name: '',
      email: '',
    } as IUser,
    first_timer: 0,
    offering: 0,
    service: '',
    children: 0,
    baby: 0,
    adult: 0,
    teen: 0,
    categories: [
      {
        traffic: '0',
      },
      {
        info_desk: '0',
      },
    ],
    created_at: new Date().toUTCString(),
  },
  isLoading: false,
  error: '',
};

export const attendanceSlice = createSlice({
  name: 'attendance',
  initialState,
  reducers: {},
});
export default attendanceSlice.reducer;
