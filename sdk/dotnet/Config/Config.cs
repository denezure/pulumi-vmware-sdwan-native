// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Xyz
{
    public static class Config
    {
        [global::System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("xyz");

        private static readonly __Value<string?> _vcoApiKey = new __Value<string?>(() => __config.Get("vcoApiKey"));
        /// <summary>
        /// API key for the VCO
        /// </summary>
        public static string? VcoApiKey
        {
            get => _vcoApiKey.Get();
            set => _vcoApiKey.Set(value);
        }

        private static readonly __Value<string?> _vcoUrl = new __Value<string?>(() => __config.Get("vcoUrl"));
        /// <summary>
        /// FQDN of the VCO
        /// </summary>
        public static string? VcoUrl
        {
            get => _vcoUrl.Get();
            set => _vcoUrl.Set(value);
        }

    }
}