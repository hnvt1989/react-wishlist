import React from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import { fetchCurrentUserWishes } from '../../actions';

class WishMyList extends React.Component {
  componentDidMount() {
    this.props.fetchCurrentUserWishes();
  }

  renderAdmin(wish) {
    if (wish.userId === this.props.currentUserId) {
      return (
        <div className="right floated content">
          <Link to={`/wishes/edit/${wish.id}`} className="ui button primary">
            Edit
          </Link>
          <Link to={`/wishes/delete/${wish.id}`} className="ui button negative">
            Delete
          </Link>
        </div>
      );
    }
  }

  renderList() {
    return this.props.currentUserWishes.map(wish => {
      return (
        <div className="item" key={wish.id}>
          {this.renderAdmin(wish)}
          <i className="large middle aligned icon gift" />
          <div className="content">
            <Link to={`/wishes/${wish.id}`} className="header">
              {wish.name}
            </Link>
            <div className="description">{wish.description}</div>
          </div>
        </div>
      );
    });
  }

  renderCreate() {
    if (this.props.isSignedIn) {
      return (
        <div style={{ textAlign: 'right' }}>
          <Link to="/wishes/new" className="ui button primary">
            Create Wish
          </Link>
        </div>
      );
    }
  }

  render() {
    return (
      <div>
        <h2>Wishes</h2>
        <div className="ui celled list">{this.renderList()}</div>
        {this.renderCreate()}
      </div>
    );
  }
}

const mapStateToProps = state => {
  return {
    currentUserWishes: Object.values(state.currentUser.wishes),
    currentUserId: state.auth.userId,
    isSignedIn: state.auth.isSignedIn
  };
};

export default connect(
  mapStateToProps,
  { fetchCurrentUserWishes }
)(WishMyList);
