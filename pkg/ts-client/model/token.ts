/* tslint:disable */
/* eslint-disable */
/**
 * HARUTO COATIONG API
 * HARUTO COATIONG API Interface
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */



/**
 * 認証の証となるトークン
 * @export
 * @interface Token
 */
export interface Token {
    /**
     * アクセストークン
     * @type {string}
     * @memberof Token
     */
    'access_token': string;
    /**
     * リフレッシュトークン
     * @type {string}
     * @memberof Token
     */
    'refresh_token': string;
    /**
     * リフレッシュトークンの有効期限(unixtimestamp utc)
     * @type {number}
     * @memberof Token
     */
    'expires_at': number;
}
