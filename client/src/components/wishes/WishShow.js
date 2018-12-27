import React from 'react';
//import flv from 'flv.js';
import { connect } from 'react-redux';
import { fetchWish } from '../../actions';

class WishShow extends React.Component {
  componentDidMount() {
    const { id } = this.props.match.params;

    this.props.fetchWish(id);
  }

  render() {
    return (
      <div>
        <h3>Wish Items</h3>
        <div className="ui celled list">{this.renderList()}</div>
      </div>
    );
  }

  renderList() {
    if (!this.props.wish) {
      return <div>Loading...</div>;
    }

    return this.props.wish['items'].map(item => {
      return (
        <div className="item" key={item.id}>
          <i className="large middle aligned icon gift" />
          <div className="content">
            <div className="header">
              <a href={`${item.url}`} target="_blank">{item.name}</a>
            </div>
            <div className="description">
              {item.note}
            </div>
          </div>
        </div>
      );
    });
  }
}

const mapStateToProps = (state, ownProps) => {
  return { wish: state.wishes[ownProps.match.params.id] };
};

export default connect(
  mapStateToProps,
  { fetchWish }
)(WishShow);
