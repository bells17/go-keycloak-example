class Account < ApplicationRecord
  def self.get_saml_settings(url_base)
    # https://github.com/SAML-Toolkits/ruby-saml#metadata-based-configuration
    idp_metadata_parser = OneLogin::RubySaml::IdpMetadataParser.new
    settings = idp_metadata_parser.parse_remote("http://localhost:8080/realms/sample/protocol/saml/descriptor")
    settings.sp_entity_id                   = "sample-client"
    settings.issuer                         = url_base + "/saml/metadata"
    settings.assertion_consumer_service_url = url_base + "/saml/acs"
    settings.assertion_consumer_logout_service_url = url_base + "/saml/logout"
    settings.protocol_binding = "urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST"
    settings.soft = true
    settings
  end
end
