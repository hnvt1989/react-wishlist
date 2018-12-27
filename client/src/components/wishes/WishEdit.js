import _ from 'lodash';
import React from 'react';
import { connect } from 'react-redux';
import { fetchWish, editWish } from '../../actions';
import WishForm from './WishForm';

class WishEdit extends React.Component {
  componentDidMount() {
    this.props.fetchWish(this.props.match.params.id);
  }

  onSubmit = formValues => {
    this.props.editWish(this.props.match.params.id, formValues);
  };

  render() {
    if (!this.props.wish) {
      return <div>Loading...</div>;
    }

    return (
      <div>
        <h3>Edit the list</h3>
        <WishForm
          initialValues={_.pick(this.props.wish, 'name', 'description', 'items')}
          onSubmit={this.onSubmit}
        />
      </div>
    );
  }
}

const mapStateToProps = (state, ownProps) => {
  return { wish: state.wishes[ownProps.match.params.id] };
};

export default connect(
  mapStateToProps,
  { fetchWish, editWish }
)(WishEdit);
