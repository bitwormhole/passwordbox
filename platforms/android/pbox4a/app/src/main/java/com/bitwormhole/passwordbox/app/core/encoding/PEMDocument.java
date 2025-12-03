package com.bitwormhole.passwordbox.app.core.encoding;

import java.util.List;

public class PEMDocument {

    private List<PEMBlock> blocks;

    public PEMDocument() {
    }

    public List<PEMBlock> getBlocks() {
        return blocks;
    }

    public void setBlocks(List<PEMBlock> blocks) {
        this.blocks = blocks;
    }
}
