import React from 'react';
import { connect } from 'react-redux';
import { createWish } from '../../actions';
import WishForm from './WishForm';

class WishCreate extends React.Component {
  onSubmit = formValues => {
    this.props.createWish(formValues);
  };

  render() {
    return (
      <div>
        <h3>Create a Wish</h3>
        <WishForm initialValues={{'name': '', 'description': '', 'items': []}} onSubmit={this.onSubmit} />
      </div>
    );
  }
}

export default connect(
  null,
  { createWish }
)(WishCreate);
