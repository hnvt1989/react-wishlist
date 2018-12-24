import React from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import Modal from '../Modal';
import history from '../../history';
import { fetchWish, deleteWish } from '../../actions';

class WishDelete extends React.Component {
  componentDidMount() {
    this.props.fetchWish(this.props.match.params.id);
  }

  renderActions() {
    const { id } = this.props.match.params;

    return (
      <React.Fragment>
        <button
          onClick={() => this.props.deleteWish(id)}
          className="ui button negative"
        >
          Delete
        </button>
        <Link to="/" className="ui button">
          Cancel
        </Link>
      </React.Fragment>
    );
  }

  renderContent() {
    if (!this.props.wish) {
      return 'Are you sure you want to delete this wish?';
    }

    return `Are you sure you want to delete the wish with title: ${
      this.props.wish.name
    }`;
  }

  render() {
    return (
      <Modal
        title="Delete Stream"
        content={this.renderContent()}
        actions={this.renderActions()}
        onDismiss={() => history.push('/')}
      />
    );
  }
}

const mapStateToProps = (state, ownProps) => {
  return { wish: state.wishes[ownProps.match.params.id] };
};

export default connect(
  mapStateToProps,
  { fetchWish, deleteWish }
)(WishDelete);
