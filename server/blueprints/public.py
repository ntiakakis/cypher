from flask import Blueprint, jsonify

public_bp = Blueprint("public_bp", __name__)

@public_bp.route('/public')
def public():
    return jsonify({
        "name": ""
    })