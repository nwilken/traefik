package label

// Traefik labels
const (
	Prefix                                                      = "traefik."
	SuffixBackend                                               = "backend"
	SuffixDomain                                                = "domain"
	SuffixEnable                                                = "enable"
	SuffixPort                                                  = "port"
	SuffixPortName                                              = "portName"
	SuffixPortIndex                                             = "portIndex"
	SuffixProtocol                                              = "protocol"
	SuffixTags                                                  = "tags"
	SuffixWeight                                                = "weight"
	SuffixBackendID                                             = "backend.id"
	SuffixBackendCircuitBreaker                                 = "backend.circuitbreaker"
	SuffixBackendCircuitBreakerExpression                       = "backend.circuitbreaker.expression"
	SuffixBackendHealthCheckScheme                              = "backend.healthcheck.scheme"
	SuffixBackendHealthCheckPath                                = "backend.healthcheck.path"
	SuffixBackendHealthCheckPort                                = "backend.healthcheck.port"
	SuffixBackendHealthCheckInterval                            = "backend.healthcheck.interval"
	SuffixBackendHealthCheckHostname                            = "backend.healthcheck.hostname"
	SuffixBackendHealthCheckHeaders                             = "backend.healthcheck.headers"
	SuffixBackendLoadBalancer                                   = "backend.loadbalancer"
	SuffixBackendLoadBalancerMethod                             = SuffixBackendLoadBalancer + ".method"
	SuffixBackendLoadBalancerSticky                             = SuffixBackendLoadBalancer + ".sticky"
	SuffixBackendLoadBalancerStickiness                         = SuffixBackendLoadBalancer + ".stickiness"
	SuffixBackendLoadBalancerStickinessCookieName               = SuffixBackendLoadBalancer + ".stickiness.cookieName"
	SuffixBackendLoadBalancerStickinessCookiePath               = SuffixBackendLoadBalancer + ".stickiness.cookiePath"
	SuffixBackendMaxConnAmount                                  = "backend.maxconn.amount"
	SuffixBackendMaxConnExtractorFunc                           = "backend.maxconn.extractorfunc"
	SuffixBackendBuffering                                      = "backend.buffering"
	SuffixBackendResponseForwardingFlushInterval                = "backend.responseForwarding.flushInterval"
	SuffixBackendBufferingMaxRequestBodyBytes                   = SuffixBackendBuffering + ".maxRequestBodyBytes"
	SuffixBackendBufferingMemRequestBodyBytes                   = SuffixBackendBuffering + ".memRequestBodyBytes"
	SuffixBackendBufferingMaxResponseBodyBytes                  = SuffixBackendBuffering + ".maxResponseBodyBytes"
	SuffixBackendBufferingMemResponseBodyBytes                  = SuffixBackendBuffering + ".memResponseBodyBytes"
	SuffixBackendBufferingRetryExpression                       = SuffixBackendBuffering + ".retryExpression"
	SuffixFrontend                                              = "frontend"
	SuffixFrontendAuth                                          = SuffixFrontend + ".auth"
	SuffixFrontendAuthBasic                                     = SuffixFrontendAuth + ".basic"
	SuffixFrontendAuthBasicRemoveHeader                         = SuffixFrontendAuthBasic + ".removeHeader"
	SuffixFrontendAuthBasicUsers                                = SuffixFrontendAuthBasic + ".users"
	SuffixFrontendAuthBasicUsersFile                            = SuffixFrontendAuthBasic + ".usersFile"
	SuffixFrontendAuthDigest                                    = SuffixFrontendAuth + ".digest"
	SuffixFrontendAuthDigestRemoveHeader                        = SuffixFrontendAuthDigest + ".removeHeader"
	SuffixFrontendAuthDigestUsers                               = SuffixFrontendAuthDigest + ".users"
	SuffixFrontendAuthDigestUsersFile                           = SuffixFrontendAuthDigest + ".usersFile"
	SuffixFrontendAuthForward                                   = SuffixFrontendAuth + ".forward"
	SuffixFrontendAuthForwardAddress                            = SuffixFrontendAuthForward + ".address"
	SuffixFrontendAuthForwardAuthResponseHeaders                = SuffixFrontendAuthForward + ".authResponseHeaders"
	SuffixFrontendAuthForwardTLS                                = SuffixFrontendAuthForward + ".tls"
	SuffixFrontendAuthForwardTLSCa                              = SuffixFrontendAuthForwardTLS + ".ca"
	SuffixFrontendAuthForwardTLSCaOptional                      = SuffixFrontendAuthForwardTLS + ".caOptional"
	SuffixFrontendAuthForwardTLSCert                            = SuffixFrontendAuthForwardTLS + ".cert"
	SuffixFrontendAuthForwardTLSInsecureSkipVerify              = SuffixFrontendAuthForwardTLS + ".insecureSkipVerify"
	SuffixFrontendAuthForwardTLSKey                             = SuffixFrontendAuthForwardTLS + ".key"
	SuffixFrontendAuthForwardTrustForwardHeader                 = SuffixFrontendAuthForward + ".trustForwardHeader"
	SuffixFrontendAuthHeaderField                               = SuffixFrontendAuth + ".headerField"
	SuffixFrontendEntryPoints                                   = "frontend.entryPoints"
	SuffixFrontendHeaders                                       = "frontend.headers."
	SuffixFrontendRequestHeaders                                = SuffixFrontendHeaders + "customRequestHeaders"
	SuffixFrontendResponseHeaders                               = SuffixFrontendHeaders + "customResponseHeaders"
	SuffixFrontendHeadersAllowedHosts                           = SuffixFrontendHeaders + "allowedHosts"
	SuffixFrontendHeadersHostsProxyHeaders                      = SuffixFrontendHeaders + "hostsProxyHeaders"
	SuffixFrontendHeadersSSLForceHost                           = SuffixFrontendHeaders + "SSLForceHost"
	SuffixFrontendHeadersSSLRedirect                            = SuffixFrontendHeaders + "SSLRedirect"
	SuffixFrontendHeadersSSLTemporaryRedirect                   = SuffixFrontendHeaders + "SSLTemporaryRedirect"
	SuffixFrontendHeadersSSLHost                                = SuffixFrontendHeaders + "SSLHost"
	SuffixFrontendHeadersSSLProxyHeaders                        = SuffixFrontendHeaders + "SSLProxyHeaders"
	SuffixFrontendHeadersSTSSeconds                             = SuffixFrontendHeaders + "STSSeconds"
	SuffixFrontendHeadersSTSIncludeSubdomains                   = SuffixFrontendHeaders + "STSIncludeSubdomains"
	SuffixFrontendHeadersSTSPreload                             = SuffixFrontendHeaders + "STSPreload"
	SuffixFrontendHeadersForceSTSHeader                         = SuffixFrontendHeaders + "forceSTSHeader"
	SuffixFrontendHeadersFrameDeny                              = SuffixFrontendHeaders + "frameDeny"
	SuffixFrontendHeadersCustomFrameOptionsValue                = SuffixFrontendHeaders + "customFrameOptionsValue"
	SuffixFrontendHeadersContentTypeNosniff                     = SuffixFrontendHeaders + "contentTypeNosniff"
	SuffixFrontendHeadersBrowserXSSFilter                       = SuffixFrontendHeaders + "browserXSSFilter"
	SuffixFrontendHeadersCustomBrowserXSSValue                  = SuffixFrontendHeaders + "customBrowserXSSValue"
	SuffixFrontendHeadersContentSecurityPolicy                  = SuffixFrontendHeaders + "contentSecurityPolicy"
	SuffixFrontendHeadersPublicKey                              = SuffixFrontendHeaders + "publicKey"
	SuffixFrontendHeadersReferrerPolicy                         = SuffixFrontendHeaders + "referrerPolicy"
	SuffixFrontendHeadersIsDevelopment                          = SuffixFrontendHeaders + "isDevelopment"
	SuffixFrontendPassHostHeader                                = "frontend.passHostHeader"
	SuffixFrontendPassTLSClientCert                             = "frontend.passTLSClientCert"
	SuffixFrontendPassTLSClientCertPem                          = SuffixFrontendPassTLSClientCert + ".pem"
	SuffixFrontendPassTLSClientCertInfos                        = SuffixFrontendPassTLSClientCert + ".infos"
	SuffixFrontendPassTLSClientCertInfosIssuer                  = SuffixFrontendPassTLSClientCertInfos + ".issuer"
	SuffixFrontendPassTLSClientCertInfosIssuerCommonName        = SuffixFrontendPassTLSClientCertInfosIssuer + ".commonName"
	SuffixFrontendPassTLSClientCertInfosIssuerCountry           = SuffixFrontendPassTLSClientCertInfosIssuer + ".country"
	SuffixFrontendPassTLSClientCertInfosIssuerDomainComponent   = SuffixFrontendPassTLSClientCertInfosIssuer + ".domainComponent"
	SuffixFrontendPassTLSClientCertInfosIssuerLocality          = SuffixFrontendPassTLSClientCertInfosIssuer + ".locality"
	SuffixFrontendPassTLSClientCertInfosIssuerOrganization      = SuffixFrontendPassTLSClientCertInfosIssuer + ".organization"
	SuffixFrontendPassTLSClientCertInfosIssuerProvince          = SuffixFrontendPassTLSClientCertInfosIssuer + ".province"
	SuffixFrontendPassTLSClientCertInfosIssuerSerialNumber      = SuffixFrontendPassTLSClientCertInfosIssuer + ".serialNumber"
	SuffixFrontendPassTLSClientCertInfosSubject                 = SuffixFrontendPassTLSClientCertInfos + ".subject"
	SuffixFrontendPassTLSClientCertInfosNotAfter                = SuffixFrontendPassTLSClientCertInfos + ".notAfter"
	SuffixFrontendPassTLSClientCertInfosNotBefore               = SuffixFrontendPassTLSClientCertInfos + ".notBefore"
	SuffixFrontendPassTLSClientCertInfosSans                    = SuffixFrontendPassTLSClientCertInfos + ".sans"
	SuffixFrontendPassTLSClientCertInfosSubjectCommonName       = SuffixFrontendPassTLSClientCertInfosSubject + ".commonName"
	SuffixFrontendPassTLSClientCertInfosSubjectCountry          = SuffixFrontendPassTLSClientCertInfosSubject + ".country"
	SuffixFrontendPassTLSClientCertInfosSubjectDomainComponent  = SuffixFrontendPassTLSClientCertInfosSubject + ".domainComponent"
	SuffixFrontendPassTLSClientCertInfosSubjectLocality         = SuffixFrontendPassTLSClientCertInfosSubject + ".locality"
	SuffixFrontendPassTLSClientCertInfosSubjectOrganization     = SuffixFrontendPassTLSClientCertInfosSubject + ".organization"
	SuffixFrontendPassTLSClientCertInfosSubjectProvince         = SuffixFrontendPassTLSClientCertInfosSubject + ".province"
	SuffixFrontendPassTLSClientCertInfosSubjectSerialNumber     = SuffixFrontendPassTLSClientCertInfosSubject + ".serialNumber"
	SuffixFrontendPassTLSCert                                   = "frontend.passTLSCert" // Deprecated
	SuffixFrontendPriority                                      = "frontend.priority"
	SuffixFrontendRateLimitExtractorFunc                        = "frontend.rateLimit.extractorFunc"
	SuffixFrontendRedirectEntryPoint                            = "frontend.redirect.entryPoint"
	SuffixFrontendRedirectRegex                                 = "frontend.redirect.regex"
	SuffixFrontendRedirectReplacement                           = "frontend.redirect.replacement"
	SuffixFrontendRedirectPermanent                             = "frontend.redirect.permanent"
	SuffixFrontendRule                                          = "frontend.rule"
	SuffixFrontendWhitelistSourceRange                          = "frontend.whitelistSourceRange" // Deprecated
	SuffixFrontendWhiteList                                     = "frontend.whiteList."
	SuffixFrontendWhiteListSourceRange                          = SuffixFrontendWhiteList + "sourceRange"
	SuffixFrontendWhiteListUseXForwardedFor                     = SuffixFrontendWhiteList + "useXForwardedFor"
	TraefikDomain                                               = Prefix + SuffixDomain
	TraefikEnable                                               = Prefix + SuffixEnable
	TraefikPort                                                 = Prefix + SuffixPort
	TraefikPortName                                             = Prefix + SuffixPortName
	TraefikPortIndex                                            = Prefix + SuffixPortIndex
	TraefikProtocol                                             = Prefix + SuffixProtocol
	TraefikTags                                                 = Prefix + SuffixTags
	TraefikWeight                                               = Prefix + SuffixWeight
	TraefikBackend                                              = Prefix + SuffixBackend
	TraefikBackendID                                            = Prefix + SuffixBackendID
	TraefikBackendCircuitBreaker                                = Prefix + SuffixBackendCircuitBreaker
	TraefikBackendCircuitBreakerExpression                      = Prefix + SuffixBackendCircuitBreakerExpression
	TraefikBackendHealthCheckScheme                             = Prefix + SuffixBackendHealthCheckScheme
	TraefikBackendHealthCheckPath                               = Prefix + SuffixBackendHealthCheckPath
	TraefikBackendHealthCheckPort                               = Prefix + SuffixBackendHealthCheckPort
	TraefikBackendHealthCheckInterval                           = Prefix + SuffixBackendHealthCheckInterval
	TraefikBackendHealthCheckHostname                           = Prefix + SuffixBackendHealthCheckHostname
	TraefikBackendHealthCheckHeaders                            = Prefix + SuffixBackendHealthCheckHeaders
	TraefikBackendLoadBalancer                                  = Prefix + SuffixBackendLoadBalancer
	TraefikBackendLoadBalancerMethod                            = Prefix + SuffixBackendLoadBalancerMethod
	TraefikBackendLoadBalancerSticky                            = Prefix + SuffixBackendLoadBalancerSticky
	TraefikBackendLoadBalancerStickiness                        = Prefix + SuffixBackendLoadBalancerStickiness
	TraefikBackendLoadBalancerStickinessCookieName              = Prefix + SuffixBackendLoadBalancerStickinessCookieName
	TraefikBackendLoadBalancerStickinessCookiePath              = Prefix + SuffixBackendLoadBalancerStickinessCookiePath
	TraefikBackendMaxConnAmount                                 = Prefix + SuffixBackendMaxConnAmount
	TraefikBackendMaxConnExtractorFunc                          = Prefix + SuffixBackendMaxConnExtractorFunc
	TraefikBackendBuffering                                     = Prefix + SuffixBackendBuffering
	TraefikBackendResponseForwardingFlushInterval               = Prefix + SuffixBackendResponseForwardingFlushInterval
	TraefikBackendBufferingMaxRequestBodyBytes                  = Prefix + SuffixBackendBufferingMaxRequestBodyBytes
	TraefikBackendBufferingMemRequestBodyBytes                  = Prefix + SuffixBackendBufferingMemRequestBodyBytes
	TraefikBackendBufferingMaxResponseBodyBytes                 = Prefix + SuffixBackendBufferingMaxResponseBodyBytes
	TraefikBackendBufferingMemResponseBodyBytes                 = Prefix + SuffixBackendBufferingMemResponseBodyBytes
	TraefikBackendBufferingRetryExpression                      = Prefix + SuffixBackendBufferingRetryExpression
	TraefikFrontend                                             = Prefix + SuffixFrontend
	TraefikFrontendAuth                                         = Prefix + SuffixFrontendAuth
	TraefikFrontendAuthBasic                                    = Prefix + SuffixFrontendAuthBasic
	TraefikFrontendAuthBasicRemoveHeader                        = Prefix + SuffixFrontendAuthBasicRemoveHeader
	TraefikFrontendAuthBasicUsers                               = Prefix + SuffixFrontendAuthBasicUsers
	TraefikFrontendAuthBasicUsersFile                           = Prefix + SuffixFrontendAuthBasicUsersFile
	TraefikFrontendAuthDigest                                   = Prefix + SuffixFrontendAuthDigest
	TraefikFrontendAuthDigestRemoveHeader                       = Prefix + SuffixFrontendAuthDigestRemoveHeader
	TraefikFrontendAuthDigestUsers                              = Prefix + SuffixFrontendAuthDigestUsers
	TraefikFrontendAuthDigestUsersFile                          = Prefix + SuffixFrontendAuthDigestUsersFile
	TraefikFrontendAuthForward                                  = Prefix + SuffixFrontendAuthForward
	TraefikFrontendAuthForwardAddress                           = Prefix + SuffixFrontendAuthForwardAddress
	TraefikFrontendAuthForwardAuthResponseHeaders               = Prefix + SuffixFrontendAuthForwardAuthResponseHeaders
	TraefikFrontendAuthForwardTLS                               = Prefix + SuffixFrontendAuthForwardTLS
	TraefikFrontendAuthForwardTLSCa                             = Prefix + SuffixFrontendAuthForwardTLSCa
	TraefikFrontendAuthForwardTLSCaOptional                     = Prefix + SuffixFrontendAuthForwardTLSCaOptional
	TraefikFrontendAuthForwardTLSCert                           = Prefix + SuffixFrontendAuthForwardTLSCert
	TraefikFrontendAuthForwardTLSInsecureSkipVerify             = Prefix + SuffixFrontendAuthForwardTLSInsecureSkipVerify
	TraefikFrontendAuthForwardTLSKey                            = Prefix + SuffixFrontendAuthForwardTLSKey
	TraefikFrontendAuthForwardTrustForwardHeader                = Prefix + SuffixFrontendAuthForwardTrustForwardHeader
	TraefikFrontendAuthHeaderField                              = Prefix + SuffixFrontendAuthHeaderField
	TraefikFrontendEntryPoints                                  = Prefix + SuffixFrontendEntryPoints
	TraefikFrontendPassHostHeader                               = Prefix + SuffixFrontendPassHostHeader
	TraefikFrontendPassTLSClientCert                            = Prefix + SuffixFrontendPassTLSClientCert
	TraefikFrontendPassTLSClientCertPem                         = Prefix + SuffixFrontendPassTLSClientCertPem
	TraefikFrontendPassTLSClientCertInfos                       = Prefix + SuffixFrontendPassTLSClientCertInfos
	TraefikFrontendPassTLSClientCertInfosIssuer                 = Prefix + SuffixFrontendPassTLSClientCertInfosIssuer
	TraefikFrontendPassTLSClientCertInfosIssuerCommonName       = Prefix + SuffixFrontendPassTLSClientCertInfosIssuerCommonName
	TraefikFrontendPassTLSClientCertInfosIssuerCountry          = Prefix + SuffixFrontendPassTLSClientCertInfosIssuerCountry
	TraefikFrontendPassTLSClientCertInfosIssuerDomainComponent  = Prefix + SuffixFrontendPassTLSClientCertInfosIssuerDomainComponent
	TraefikFrontendPassTLSClientCertInfosIssuerLocality         = Prefix + SuffixFrontendPassTLSClientCertInfosIssuerLocality
	TraefikFrontendPassTLSClientCertInfosIssuerOrganization     = Prefix + SuffixFrontendPassTLSClientCertInfosIssuerOrganization
	TraefikFrontendPassTLSClientCertInfosIssuerProvince         = Prefix + SuffixFrontendPassTLSClientCertInfosIssuerProvince
	TraefikFrontendPassTLSClientCertInfosIssuerSerialNumber     = Prefix + SuffixFrontendPassTLSClientCertInfosIssuerSerialNumber
	TraefikFrontendPassTLSClientCertInfosNotAfter               = Prefix + SuffixFrontendPassTLSClientCertInfosNotAfter
	TraefikFrontendPassTLSClientCertInfosNotBefore              = Prefix + SuffixFrontendPassTLSClientCertInfosNotBefore
	TraefikFrontendPassTLSClientCertInfosSans                   = Prefix + SuffixFrontendPassTLSClientCertInfosSans
	TraefikFrontendPassTLSClientCertInfosSubject                = Prefix + SuffixFrontendPassTLSClientCertInfosSubject
	TraefikFrontendPassTLSClientCertInfosSubjectCommonName      = Prefix + SuffixFrontendPassTLSClientCertInfosSubjectCommonName
	TraefikFrontendPassTLSClientCertInfosSubjectCountry         = Prefix + SuffixFrontendPassTLSClientCertInfosSubjectCountry
	TraefikFrontendPassTLSClientCertInfosSubjectDomainComponent = Prefix + SuffixFrontendPassTLSClientCertInfosSubjectDomainComponent
	TraefikFrontendPassTLSClientCertInfosSubjectLocality        = Prefix + SuffixFrontendPassTLSClientCertInfosSubjectLocality
	TraefikFrontendPassTLSClientCertInfosSubjectOrganization    = Prefix + SuffixFrontendPassTLSClientCertInfosSubjectOrganization
	TraefikFrontendPassTLSClientCertInfosSubjectProvince        = Prefix + SuffixFrontendPassTLSClientCertInfosSubjectProvince
	TraefikFrontendPassTLSClientCertInfosSubjectSerialNumber    = Prefix + SuffixFrontendPassTLSClientCertInfosSubjectSerialNumber
	TraefikFrontendPassTLSCert                                  = Prefix + SuffixFrontendPassTLSCert // Deprecated
	TraefikFrontendPriority                                     = Prefix + SuffixFrontendPriority
	TraefikFrontendRateLimitExtractorFunc                       = Prefix + SuffixFrontendRateLimitExtractorFunc
	TraefikFrontendRedirectEntryPoint                           = Prefix + SuffixFrontendRedirectEntryPoint
	TraefikFrontendRedirectRegex                                = Prefix + SuffixFrontendRedirectRegex
	TraefikFrontendRedirectReplacement                          = Prefix + SuffixFrontendRedirectReplacement
	TraefikFrontendRedirectPermanent                            = Prefix + SuffixFrontendRedirectPermanent
	TraefikFrontendRule                                         = Prefix + SuffixFrontendRule
	TraefikFrontendWhitelistSourceRange                         = Prefix + SuffixFrontendWhitelistSourceRange // Deprecated
	TraefikFrontendWhiteListSourceRange                         = Prefix + SuffixFrontendWhiteListSourceRange
	TraefikFrontendWhiteListUseXForwardedFor                    = Prefix + SuffixFrontendWhiteListUseXForwardedFor
	TraefikFrontendRequestHeaders                               = Prefix + SuffixFrontendRequestHeaders
	TraefikFrontendResponseHeaders                              = Prefix + SuffixFrontendResponseHeaders
	TraefikFrontendAllowedHosts                                 = Prefix + SuffixFrontendHeadersAllowedHosts
	TraefikFrontendHostsProxyHeaders                            = Prefix + SuffixFrontendHeadersHostsProxyHeaders
	TraefikFrontendSSLForceHost                                 = Prefix + SuffixFrontendHeadersSSLForceHost
	TraefikFrontendSSLRedirect                                  = Prefix + SuffixFrontendHeadersSSLRedirect
	TraefikFrontendSSLTemporaryRedirect                         = Prefix + SuffixFrontendHeadersSSLTemporaryRedirect
	TraefikFrontendSSLHost                                      = Prefix + SuffixFrontendHeadersSSLHost
	TraefikFrontendSSLProxyHeaders                              = Prefix + SuffixFrontendHeadersSSLProxyHeaders
	TraefikFrontendSTSSeconds                                   = Prefix + SuffixFrontendHeadersSTSSeconds
	TraefikFrontendSTSIncludeSubdomains                         = Prefix + SuffixFrontendHeadersSTSIncludeSubdomains
	TraefikFrontendSTSPreload                                   = Prefix + SuffixFrontendHeadersSTSPreload
	TraefikFrontendForceSTSHeader                               = Prefix + SuffixFrontendHeadersForceSTSHeader
	TraefikFrontendFrameDeny                                    = Prefix + SuffixFrontendHeadersFrameDeny
	TraefikFrontendCustomFrameOptionsValue                      = Prefix + SuffixFrontendHeadersCustomFrameOptionsValue
	TraefikFrontendContentTypeNosniff                           = Prefix + SuffixFrontendHeadersContentTypeNosniff
	TraefikFrontendBrowserXSSFilter                             = Prefix + SuffixFrontendHeadersBrowserXSSFilter
	TraefikFrontendCustomBrowserXSSValue                        = Prefix + SuffixFrontendHeadersCustomBrowserXSSValue
	TraefikFrontendContentSecurityPolicy                        = Prefix + SuffixFrontendHeadersContentSecurityPolicy
	TraefikFrontendPublicKey                                    = Prefix + SuffixFrontendHeadersPublicKey
	TraefikFrontendReferrerPolicy                               = Prefix + SuffixFrontendHeadersReferrerPolicy
	TraefikFrontendIsDevelopment                                = Prefix + SuffixFrontendHeadersIsDevelopment
	BaseFrontendErrorPage                                       = "frontend.errors."
	SuffixErrorPageBackend                                      = "backend"
	SuffixErrorPageQuery                                        = "query"
	SuffixErrorPageStatus                                       = "status"
	BaseFrontendRateLimit                                       = "frontend.rateLimit.rateSet."
	SuffixRateLimitPeriod                                       = "period"
	SuffixRateLimitAverage                                      = "average"
	SuffixRateLimitBurst                                        = "burst"
)
