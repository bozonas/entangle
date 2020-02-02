import React, { useEffect } from "react";
import { Form, Button } from "semantic-ui-react";
import { useForm } from "react-hook-form";
import crypto from 'crypto-js';

function getRandomString(length) {
  var randomChars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_!';
  var result = '';
  for ( var i = 0; i < length; i++ ) {
      result += randomChars.charAt(Math.floor(Math.random() * randomChars.length));
  }
  return result;
}

const FormExampleFieldError = () => {
  useEffect(() => {
    register({ name: "secret" }, { required: true });
  }, []);

  const { register, errors, handleSubmit, setValue, triggerValidation } = useForm()
  const onSubmit = data => {
    let key: string = getRandomString(20);
    console.log(key);
    var ciphertext = crypto.AES.encrypt(JSON.stringify(data), key).toString();
    console.log(ciphertext);

    // send here!

    console.log(data);
  }

  return (
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
  );
};

export default FormExampleFieldError;