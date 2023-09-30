import { combineReducers } from 'redux';

import userReducer from './user';
import attendanceReduer from './attendance';

const rootReducer = combineReducers({
  account: combineReducers({
    user: userReducer,
    attendance: attendanceReduer,
  }),
});

export default rootReducer;
