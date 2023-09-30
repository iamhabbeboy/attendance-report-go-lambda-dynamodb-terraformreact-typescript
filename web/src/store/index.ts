import { configureStore, ThunkAction, Action } from '@reduxjs/toolkit';
import logger from 'redux-logger';
// import storageSession from 'redux-persist/lib/storage/session';
// import { persistStore, persistReducer } from 'redux-persist';
import rootReducer from './collections';
// import { checkUserAccountType } from '../middleware/checkUserAccountType';

// const persistConfig = {
//   key: 'root',
//   version: 1,
//   storage: storageSession,
//   // blacklist: [""],
// };
// const persistedReducer = persistReducer(persistConfig, rootReducer);

const store = configureStore({
  reducer: rootReducer,
  middleware: (getDefaultMiddleware) =>
    process.env.NODE_ENV === 'development'
      ? getDefaultMiddleware({ serializableCheck: false }).concat(logger)
      : getDefaultMiddleware({ serializableCheck: false }),
  devTools: process.env.NODE_ENV === 'development',
});
// const persistor = persistStore(store);
export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;
export { store };
