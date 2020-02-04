import React, { useEffect, useState } from "react";
import { Form, Button, Grid, Transition, Input, Dimmer, Loader } from "semantic-ui-react";
import { SemanticToastContainer, toast } from 'react-semantic-toasts';
import { useForm } from "react-hook-form";
import crypto from 'crypto-js';
import axios from 'axios';
import { API_URL } from '../constants';

function getRandomString(length) {
  var randomChars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz';
  var result = '';
  for (var i = 0; i < length; i++) {
    result += randomChars.charAt(Math.floor(Math.random() * randomChars.length));
  }
  return result;
}

interface Request {
  key: string
  ciphertext: string
}

interface Response { }

const MainForm = (props) => {
  useEffect(() => {
    register({ name: "secret" }, { required: true });
  }, []);

  const [url, setUrl] = useState("");
  const [formLoading, setLoading] = useState(false);
  const [visibility, setVisibility] = useState({ form: true, copy: false });

  const { register, errors, handleSubmit, setValue, triggerValidation } = useForm()
  const onSubmit = async data => {
    setLoading(true);
    const key: string = getRandomString(10);
    const pass: string = getRandomString(20);
    const ciphertext = crypto.AES.encrypt(data.secret, pass).toString();

    const request: Request = {
      key: key,
      ciphertext: ciphertext
    }
    let response: Response
    try {
      response = await axios.post<Request, Response>(`${API_URL}/message`, request);
      console.log(response);
      setUrl(`${window.location.href}${key}-${pass}`);
      setVisibility({ form: false, copy: false });
      setTimeout(function(){ setVisibility({ form: false, copy: true }); }, 500);
    } catch (e) {
      props.history.push('/errorPage');
    }
    setLoading(false);
  }

  const onIconCopy = () => {
    navigator.clipboard.writeText(url);
    toast({
      type: 'success',
      icon: 'copy',
      title: 'Copied',
      description: 'URL has been copied to your clipboard',
      animation: 'fade up',
      time: 3000,
    });
  }

  return (
    <>
      <Grid>
        <Grid.Row centered>
          <Grid.Column width={8} textAlign="center">
            <Transition visible={visibility.copy} animation='fade up' duration={500} de>
              <>
                <Input fluid
                  defaultValue={url}
                  action={{
                    color: 'teal',
                    labelPosition: 'right',
                    icon: 'copy',
                    content: 'Copy',
                    onClick: onIconCopy
                  }} />
              </>
            </Transition>
            <Transition visible={visibility.form} animation='fade up' duration={500}>
              <Form onSubmit={handleSubmit(onSubmit)}>
                <Form.TextArea
                  rows={5}
                  name="secret"
                  placeholder='Enter your secret'
                  onChange={async (e, { name, value }) => {
                    setValue(name, value);
                    await triggerValidation(name);
                  }}
                  error={errors.secret ? true : false}
                />
                <Button>Save secret</Button>
              </Form>
            </Transition>
          </Grid.Column>
        </Grid.Row>
      </Grid>
      <SemanticToastContainer position="bottom-right" />
      <Dimmer
        active={formLoading}
        page={true}>
        <Loader inverted>Loading</Loader>
      </Dimmer>
    </>
  );
};

export default MainForm;