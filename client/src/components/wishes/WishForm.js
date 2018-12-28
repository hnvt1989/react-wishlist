import React from 'react';
import { Field, FieldArray, reduxForm } from 'redux-form';

class WishForm extends React.Component {
  renderError({ error, touched }) {
    if (touched && error) {
      return (
        <div className="ui error message">
          <div className="header">{error}</div>
        </div>
      );
    }
  }

  renderInput = ({ input, label, meta }) => {
    const className = `field ${meta.error && meta.touched ? 'error' : ''}`;
    return (
      <div className={className}>
        <label>{label}</label>
        <input {...input} autoComplete="off" />
        {this.renderError(meta)}
      </div>
    );
  };

  onSubmit = formValues => {
    this.props.onSubmit(formValues);
  };

  render() {
    return (
      <form
        onSubmit={this.props.handleSubmit(this.onSubmit)}
        className="ui form error"
      >
        <Field name="name" component={this.renderInput} label="Enter Name" />
        <Field
          name="description"
          component={this.renderInput}
          label="Enter Description"
        />
        <FieldArray name="items" component={this.renderItems} label="Items" />
        <div>
          <button className="ui button primary">Submit</button>
        </div>
      </form>
    );
  }

  renderItems = ({ fields, meta: { error, submitFailed } }) => {

    if (!fields) {
      return <div>Loading...</div>;
    }

    return (
      <div>
        <div className="ui button" onClick={() => fields.push()}>
          Add Item
        </div>
        <br />
        {fields.map((field, index) => (
          <div key={index} className="segment">
            <div className="ui segment">
              <h5>Item #{index + 1}</h5>
              <Field
                name={`${field}.name`}
                type="text"
                component={this.renderInput}
                label="Item name"
              />
              <Field
                name={`${field}.url`}
                type="text"
                component={this.renderInput}
                label="Url"
              />
              <Field
                name={`${field}.note`}
                type="text"
                component={this.renderInput}
                label="Note"
              />
              <div className="ui button" title="Remove Item" onClick={() => fields.remove(index)}>
                Remove item
              </div>
            </div>
          </div>
        ))}
      </div>
    );

  }
}

const validate = formValues => {
  const errors = {};

  if (!formValues.name) {
    errors.name = 'You must enter a title';
  }

  if (!formValues.description) {
    errors.description = 'You must enter a description';
  }

  return errors;
};

export default reduxForm({
  form: 'wishForm',
  validate,
  // fields: ['name', 'description', 'fields']
})(WishForm);
