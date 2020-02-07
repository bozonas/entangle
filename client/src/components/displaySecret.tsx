import React, { useEffect, useState } from "react";
import { Grid, Form, TextArea, Dimmer, Loader } from "semantic-ui-react";
import axios from 'axios';
import crypto from 'crypto-js';
import { API_URL } from '../constants';

const DisplaySecret = (props) => {
    const [data, setData] = useState("");
    const [formLoading, setLoading] = useState(true);
    useEffect(() => {
        setLoading(true);
        const [key, pass] = props.match.params.id.split('-');
        axios(
            `${API_URL}/message/${key}`,
        ).then(res => {
            const ciphertext = res.data.ciphertext;
            const secret = crypto.AES.decrypt(ciphertext, pass).toString(crypto.enc.Utf8);
            setData(secret);
            setLoading(false);
        }).catch(err => {
            props.history.push(`/notFound`);
        });
    }, []);
    return (
        <>
            <Grid>
                <Grid.Row centered>
                    <Grid.Column width={8} textAlign="center">
                        <Form>
                            <TextArea
                                rows={5}
                                name="secret"
                                defaultValue={data}
                            />
                        </Form>
                    </Grid.Column>
                </Grid.Row>
            </Grid>
            <Dimmer
                active={formLoading}
                page={true}>
                <Loader inverted>Loading</Loader>
            </Dimmer>
        </>
    );
};

export default DisplaySecret;