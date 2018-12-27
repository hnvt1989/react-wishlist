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
    //console.log(this.props.items);
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
        <FieldArray name="items" component={this.renderItems} />
        <div>
          <button className="ui button primary">Submit</button>
        </div>
      </form>
    );
  }

  renderItems = ({ items, meta: { error, submitFailed } }) => {
    return (
      <ul>
        <li>
          {/* <button type="button" onClick={() => items.push()}>
            Add Item
      </button> */}
        </li>
        {items.map((item, index) => (
          <li key={index}>
            <button
              type="button"
              title="Remove Item"
              onClick={() => items.remove(index)}
            />

            <h4>Member #{index + 1}</h4>
            <Field
              name={`${item}.name`}
              type="text"
              component={this.renderInput}
              label="Item name"
            />
            <Field
              name={`${item}.url`}
              type="text"
              component={this.renderInput}
              label="Url"
            />

            {/* <div className="ui segment">
            <Field name={item.name} component={this.renderInput} label="Name" />
          </div>
          <div className="ui segment">
            <Field name={item.note} component={this.renderInput} label="Note" />
          </div>
          <div className="ui segment">
            <Field name={item.url} component={this.renderInput} label="Url" />
          </div> */}
          </li>
        ))}
      </ul>
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
  initialValues: {
    name: 'hey',
    description: 'sup',
    items: [{id: 1, name: 'john', description: 'Doe', url: 'hey' }]
  }
})(WishForm);
