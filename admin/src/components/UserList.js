import * as React from 'react';
import {
    List,
    Datagrid,
    TextField,
    EmailField
} from 'react-admin';

export const UserList = () => (
    <List>
        <Datagrid rowClick="edit">
            <TextField source="id" />
            <TextField source="alexaId" />
            <TextField source="alexaDevice" />
            <EmailField source="email" />
            <TextField source="lastUpdated" />
        </Datagrid>
    </List>
);