package com.bitwormhole.passwordbox.app.core.encoding;

import java.util.Map;

public final class PEMBlock {

    private String type;
    private Map<String, String> headers;
    private byte[] content;

    public PEMBlock() {
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public Map<String, String> getHeaders() {
        return headers;
    }

    public void setHeaders(Map<String, String> headers) {
        this.headers = headers;
    }

    public byte[] getContent() {
        return content;
    }

    public void setContent(byte[] content) {
        this.content = content;
    }
}
