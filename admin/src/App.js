import * as React from 'react';
import { Admin, Resource } from 'react-admin';
import { AnswerList } from './components/AnswerList';
import { UserList } from './components/UserList';
import dataProvider from './providers/dataProvider';


const App = () => (
  <Admin dataProvider={dataProvider}>
    <Resource name="users" list={UserList}/>
    <Resource name="answers" list={AnswerList} />
  </Admin>
)

export default App;
