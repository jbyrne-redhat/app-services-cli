"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[239],{3905:function(e,t,r){r.d(t,{Zo:function(){return f},kt:function(){return d}});var a=r(7294);function i(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function n(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,a)}return r}function o(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?n(Object(r),!0).forEach((function(t){i(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):n(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function s(e,t){if(null==e)return{};var r,a,i=function(e,t){if(null==e)return{};var r,a,i={},n=Object.keys(e);for(a=0;a<n.length;a++)r=n[a],t.indexOf(r)>=0||(i[r]=e[r]);return i}(e,t);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);for(a=0;a<n.length;a++)r=n[a],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(i[r]=e[r])}return i}var c=a.createContext({}),p=function(e){var t=a.useContext(c),r=t;return e&&(r="function"==typeof e?e(t):o(o({},t),e)),r},f=function(e){var t=p(e.components);return a.createElement(c.Provider,{value:t},e.children)},l={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},u=a.forwardRef((function(e,t){var r=e.components,i=e.mdxType,n=e.originalType,c=e.parentName,f=s(e,["components","mdxType","originalType","parentName"]),u=p(r),d=i,m=u["".concat(c,".").concat(d)]||u[d]||l[d]||n;return r?a.createElement(m,o(o({ref:t},f),{},{components:r})):a.createElement(m,o({ref:t},f))}));function d(e,t){var r=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var n=r.length,o=new Array(n);o[0]=u;var s={};for(var c in t)hasOwnProperty.call(t,c)&&(s[c]=t[c]);s.originalType=e,s.mdxType="string"==typeof e?e:i,o[1]=s;for(var p=2;p<n;p++)o[p]=r[p];return a.createElement.apply(null,o)}return a.createElement.apply(null,r)}u.displayName="MDXCreateElement"},1318:function(e,t,r){r.r(t),r.d(t,{frontMatter:function(){return s},contentTitle:function(){return c},metadata:function(){return p},toc:function(){return f},default:function(){return u}});var a=r(7462),i=r(3366),n=(r(7294),r(3905)),o=["components"],s={},c=void 0,p={unversionedId:"commands/rhoas_service-registry_artifact_get",id:"commands/rhoas_service-registry_artifact_get",isDocsHomePage:!1,title:"rhoas_service-registry_artifact_get",description:"rhoas service-registry artifact get",source:"@site/docs/commands/rhoas_service-registry_artifact_get.md",sourceDirName:"commands",slug:"/commands/rhoas_service-registry_artifact_get",permalink:"/app-services-cli/commands/rhoas_service-registry_artifact_get",editUrl:"https://github.com/redhat-developer/app-services-cli/docs/commands/rhoas_service-registry_artifact_get.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"rhoas_service-registry_artifact_export",permalink:"/app-services-cli/commands/rhoas_service-registry_artifact_export"},next:{title:"rhoas_service-registry_artifact_import",permalink:"/app-services-cli/commands/rhoas_service-registry_artifact_import"}},f=[{value:"rhoas service-registry artifact get",id:"rhoas-service-registry-artifact-get",children:[{value:"Synopsis",id:"synopsis",children:[]},{value:"Examples",id:"examples",children:[]},{value:"Options",id:"options",children:[]},{value:"Options inherited from parent commands",id:"options-inherited-from-parent-commands",children:[]},{value:"SEE ALSO",id:"see-also",children:[]}]}],l={toc:f};function u(e){var t=e.components,r=(0,i.Z)(e,o);return(0,n.kt)("wrapper",(0,a.Z)({},l,r,{components:t,mdxType:"MDXLayout"}),(0,n.kt)("h2",{id:"rhoas-service-registry-artifact-get"},"rhoas service-registry artifact get"),(0,n.kt)("p",null,"Get artifact by ID and group"),(0,n.kt)("h3",{id:"synopsis"},"Synopsis"),(0,n.kt)("p",null,"Get artifact by specifying ID and group.\nCommand fetches the latest artifact from the registry based on the artifact-id and group."),(0,n.kt)("p",null,'When --version is specified, the command fetches the specific artifact version.\nGet command fetches artifacts based on --group and --artifact-id and --version.\nFor fetching artifacts using global identifiers, use the "service-registry download" command'),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre"},"rhoas service-registry artifact get [flags]\n")),(0,n.kt)("h3",{id:"examples"},"Examples"),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre"},'## Get latest artifact with name "my-artifact" and print it out to standard out\nrhoas service-registry artifact get --artifact-id=my-artifact\n\n## Get latest artifact with name "my-artifact" from group "my-group" and save it to artifact.json file\nrhoas service-registry artifact get --artifact-id=my-artifact --group=my-group --output-file=artifact.json\n\n## Get latest artifact and pipe it to another command\nrhoas service-registry artifact get --artifact-id=my-artifact | grep -i \'user\'\n\n## Get artifact with custom version and print it out to standard out\nrhoas service-registry artifact get --artifact-id=myartifact --version=4\n\n')),(0,n.kt)("h3",{id:"options"},"Options"),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre"},'      --artifact-id string   ID of the artifact\n  -g, --group string         Artifact group (default "default")\n      --instance-id string   ID of the Service Registry instance to be used. By default, uses the currently selected instance\n      --output-file string   Location of the output file\n      --version string       Version of the artifact\n')),(0,n.kt)("h3",{id:"options-inherited-from-parent-commands"},"Options inherited from parent commands"),(0,n.kt)("pre",null,(0,n.kt)("code",{parentName:"pre"},"  -h, --help      Show help for a command\n  -v, --verbose   Enable verbose mode\n")),(0,n.kt)("h3",{id:"see-also"},"SEE ALSO"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},(0,n.kt)("a",{parentName:"li",href:"/app-services-cli/commands/rhoas_service-registry_artifact"},"rhoas service-registry artifact"),"\t - Manage Service Registry artifacts")))}u.isMDXComponent=!0}}]);