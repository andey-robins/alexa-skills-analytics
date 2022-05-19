import * as React from 'react';
import {
    List,
    Datagrid,
    TextField,
} from 'react-admin';

export const AnswerList = () => (
    <List>
        <Datagrid rowClick="edit">
            <TextField source="id" />
            <TextField source="questionId" />
            <TextField source="answer" />
            <TextField source="time" />
        </Datagrid>
    </List>
);