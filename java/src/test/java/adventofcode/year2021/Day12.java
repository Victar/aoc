package adventofcode.year2021;

import java.util.*;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;

public class Day12 extends BaseTest {

	@Test public void runSilver() throws Exception {
		run(false);
	}

	@Test public void runGold() throws Exception {
		run(true);
	}

	public void run(final boolean firstVisit) throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day12/input.txt");
		final Map<String, Path> allMap = new HashMap<>();
		for (final String input : data) {
			final String[] arr = StringUtils.split(input, "-");
			final Path from;
			final Path to;
			if (allMap.containsKey(arr[0])) {
				from = allMap.get(arr[0]);
			} else {
				from = new Path(arr[0]);
				allMap.put(arr[0], from);
			}
			if (allMap.containsKey(arr[1])) {
				to = allMap.get(arr[1]);
			} else {
				to = new Path(arr[1]);
				allMap.put(arr[1], to);
			}
			from.addTo(to);
			to.addTo(from);
		}

		final Path start = allMap.get("start");
		final List<Path> visitedPaths = new ArrayList<>();
		visitedPaths.add(start);
		final int count = getRoutesCount(visitedPaths, start, firstVisit);
		System.out.println(count);
	}

	public int getRoutesCount(final List<Path> route, final Path current, final boolean firstVisit) {
		if (current.isEnd()) {
			return 1;
		}
		int childCount = 0;
		for (final Path child : current.getNextPaths()) {
			if (child.isBig()) {
				route.add(child);
				childCount += getRoutesCount(route, child, firstVisit);
				route.remove(route.size() - 1);
			} else if (!route.contains(child)) {
				route.add(child);
				childCount += getRoutesCount(route, child, firstVisit);
				route.remove(route.size() - 1);
			} else if (firstVisit && !child.isEnd() && !child.isStart()) {
				route.add(child);
				childCount += getRoutesCount(route, child, false);
				route.remove(route.size() - 1);
			}
		}
		return childCount;
	}

	@Data public static class Path {

		private String name;
		private boolean big;

		private Set<Path> nextPaths = new HashSet<>();

		public Path(final String name) {
			this.name = name;
			this.big = name.equals(name.toUpperCase());
		}

		public boolean isEnd() {
			return this.name.equals("end");
		}

		public boolean isStart() {
			return this.name.equals("start");
		}

		public void addTo(final Path to) {
			this.nextPaths.add(to);
		}

		@Override public boolean equals(final Object o) {
			if (this == o) return true;
			if (o == null || getClass() != o.getClass()) return false;

			final Path path = (Path) o;

			if (this.big != path.big) return false;
			return this.name != null ? this.name.equals(path.name) : path.name == null;
		}

		@Override public int hashCode() {
			int result = this.name != null ? this.name.hashCode() : 0;
			result = 31 * result + (this.big ? 1 : 0);
			return result;
		}

		@Override public String toString() {
			return "Path{" + "name='" + this.name + '\'' + ", big=" + this.big + '}';
		}
	}

}
