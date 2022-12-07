package adventofcode.year2022;

import adventofcode.BaseTest;
import lombok.Data;
import org.junit.Ignore;
import org.junit.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Day7 extends BaseTest {

	public static final int DAY = 7;

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test public void runSilver() throws Exception {
		Dir root = readRoot();

		List<Dir> dirs = new ArrayList<>();
		root.collectSilver(dirs, 100000);
		int ans = 0;
		for (Dir d : dirs) {
			ans += d.getSize();
		}
		System.out.println(ans);
	}

	@Test public void runGold() throws Exception {
		Dir root = readRoot();
		int total = 70000000;
		int needSpace = 30000000;
		int toDelete = needSpace - (total - root.getSize());

		List<Dir> dirs = new ArrayList<>();
		root.collectGold(dirs, toDelete);
		int ans = Integer.MAX_VALUE;
		for (Dir d : dirs) {
			ans = Math.min(ans, d.getSize());
		}
		System.out.println(ans);

	}

	Dir readRoot() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day7/input.txt");
		Dir root = new Dir("/", null);
		Dir currentDir = root;
		boolean currentLS = false;
		for (int i = 1; i < data.size(); i++) {
			String input = data.get(i);
			if (currentLS && !input.startsWith("$")) {
				currentDir.addFileOrFolder(input);
			}
			if (input.startsWith("$ ls")) {
				currentLS = true;
			}
			if (input.startsWith("$ cd")) {
				currentLS = false;
				currentDir = currentDir.getDirByName(input);
			}
		}
		return root;
	}

	@Data class Dir {

		String name;
		Dir parent;
		List<Dir> subdirs = new ArrayList<>();
		Map<String, Integer> files = new HashMap<>();

		public Dir(String name, Dir parent) {
			this.name = name;
			this.parent = parent;
		}

		public void addFileOrFolder(String input) {
			String[] parts = input.split(" ");
			if ("dir".equals(parts[0])) {
				Dir subdir = new Dir(parts[1], this);
				this.subdirs.add(subdir);
			} else {
				this.files.put(parts[1], Integer.parseInt(parts[0]));
			}
		}

		public Dir getDirByName(String input) {
			String dirName = input.replaceAll("\\$ cd ", "");
			if ("..".equals(dirName)) {
				return parent;
			} else {
				for (Dir subdir : subdirs) {
					if (subdir.name.equals(dirName)) {
						return subdir;
					}
				}
			}
			return null;
		}

		public void collectSilver(List<Dir> dirsList, int size) {
			int curSize = getSize();
			if (curSize < size) {
				dirsList.add(this);
			}
			for (Dir subdir : this.subdirs) {
				subdir.collectSilver(dirsList, size);
			}
		}

		public void collectGold(List<Dir> dirsList, int size) {
			int curSize = getSize();
			if (curSize >= size) {
				dirsList.add(this);
			}
			for (Dir subdir : this.subdirs) {
				subdir.collectGold(dirsList, size);
			}
		}

		public int getSize() {
			int size = 0;

			for (Map.Entry<String, Integer> entry : files.entrySet()) {
				size += entry.getValue();
			}
			for (Dir subdir : subdirs) {
				size += subdir.getSize();
			}
			return size;
		}
	}

}
