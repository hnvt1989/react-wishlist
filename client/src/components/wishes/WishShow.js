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
        <h3>Info</h3>
        {this.renderListInfo()}
        <h3>Items</h3>
        <div className="ui celled list">{this.renderList()}</div>
      </div>
    );
  }

  renderListInfo(){
    if (!this.props.wish) {
      return <div>Loading...</div>;
    }

    return(
      <div>
        <div>{this.props.wish.name}</div>
        <div>{this.props.wish.description}</div>
      </div>
    );

  }
  renderList() {
    if (!this.props.wish) {
      return <div>Loading...</div>;
    }

    if(this.props.wish.items.length > 0){
      return this.props.wish.items.map(item => {
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
    }else{
      return <div>Empty list</div>
    }

  }
}

const mapStateToProps = (state, ownProps) => {
  return { wish: state.wishes[ownProps.match.params.id] };
};

export default connect(
  mapStateToProps,
  { fetchWish }
)(WishShow);
